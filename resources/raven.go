package resources

import (
	"github.com/getsentry/raven-go"
)

func NewRavenGo(env *Environment) (*raven.Client, error) {
	return raven.New(env.SentryDSN)
}
