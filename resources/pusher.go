package resources

import (
	"github.com/pusher/pusher-http-go"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PusherClient interface {
	AuthenticatePrivateChannel([]byte) ([]byte, error)
	Trigger(string, string, interface{}) (*pusher.BufferedEvents, error)
}

func NewPusherClient(env *Environment, httpClient *http.Client, logger *logrus.Logger) (PusherClient, error) {
	client, err := pusher.ClientFromURL(env.PusherUrl)
	if err != nil {
		logger.WithError(err).Error("incorrect $PUSHER_URL")
		return nil, err
	}
	client.HttpClient = httpClient
	return client, nil
}
