package server

import (
	"github.com/labstack/echo"
	"github.com/mikefaraponov/pentestable-auther/resources"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type (
	Server interface {
		AuthenticateWeb(echo.Context) error
		AuthenticateMobile(echo.Context) error
		RenderHTML(echo.Context) error
	}
	server struct {
		resources.PusherClient
		*resources.Environment
		*logrus.Logger
	}
)

func (s *server) AuthenticateWeb(ctx echo.Context) error {
	defer s.Info("web authenticated")
	callback := ctx.QueryParam(s.Callback)
	params := []byte(ctx.QueryString())
	response, err := s.AuthenticatePrivateChannel(params)
	if err != nil {
		s.WithError(err).Error("failed to authenticate pusher params")
		return err
	}
	return ctx.JSONPBlob(http.StatusOK, callback, response)
}

func (s *server) AuthenticateMobile(ctx echo.Context) error {
	defer s.Info("mobile authenticated")
	request := ctx.Request()
	params, err := ioutil.ReadAll(request.Body)
	if err != nil {
		s.WithError(err).Error("failed to read pusher params from request")
		return err
	}
	response, err := s.AuthenticatePrivateChannel(params)
	if err == nil {
		s.WithError(err).Error("failed to authenticate pusher params")
		return err
	}
	return ctx.JSON(http.StatusOK, response)
}

func (s *server) RenderHTML(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, HTML)
}

func New(env *resources.Environment, pusher resources.PusherClient, logger *logrus.Logger) Server {
	return &server{
		Environment:  env,
		PusherClient: pusher,
		Logger:       logger,
	}
}

const HTML = `
<!DOCTYPE html>
<head>
  <title>Pusher Test</title>
  <script src="https://js.pusher.com/4.3/pusher.min.js"></script>
  <script>
    Pusher.logToConsole = true;

    var pusher = new Pusher('72db221f9b47b4489c50', {
      cluster: 'eu',
      forceTLS: true,
      authEndpoint: 'https://d552a6d0.ngrok.io/pusher_auth_mobile',
      // authTransport: 'jsonp',
    });

    var channel = pusher.subscribe('private-my-channel');
    channel.bind('my-event', function(data) {
      alert(JSON.stringify(data));
    });
  </script>
</head>
<body>
  <h1>Pusher Test</h1>
  <p>
    Try publishing an event to channel <code>my-channel</code>
    with event name <code>my-event</code>.
  </p>
</body>
`
