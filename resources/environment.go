package resources

import "github.com/caarlos0/env"

type Environment struct {
	SentryDSN string `env:"SENTRY_DSN,required"`
	Address   string `env:"ADDRESS,required"`
	PusherUrl string `env:"PUSHER_URL,required"`
	Callback  string `env:"CALLBACK" envDefault:"callback"`
}

func NewEnvironment() (*Environment, error) {
	cfg := new(Environment)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
