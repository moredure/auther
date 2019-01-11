package bootstrap

import (
	"context"
	"github.com/labstack/echo"
	"github.com/mikefaraponov/auther/resources"
	"github.com/mikefaraponov/auther/server"
	"go.uber.org/fx"
)

type BootstrapOptions struct {
	fx.In
	fx.Lifecycle
	*echo.Echo
	server.Server
	*resources.Environment
}

func Invoke(app BootstrapOptions) {
	app.GET("/pusher_auth_web", app.AuthenticateWeb)
	app.POST("/pusher_auth_mobile", app.AuthenticateMobile)
	app.GET("/", app.RenderHTML)

	app.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := app.Start(app.Address); err != nil {
					app.Logger.Info("shutting down the server")
				}
			}()
			return nil
		},
		OnStop: func(c context.Context) error {
			return app.Shutdown(c)
		},
	})
}
