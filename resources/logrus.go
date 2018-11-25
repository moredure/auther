package resources

import (
	"github.com/getsentry/raven-go"
	"github.com/pentestable/logrus_sentry"
	"github.com/sirupsen/logrus"
)

func NewLogrusErrorLevel() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}

func NewSentryHook(raven *raven.Client, levels []logrus.Level) (*logrus_sentry.SentryHook, error) {
	return logrus_sentry.NewAsyncWithClientSentryHook(raven, levels)
}

func NewLogrusLogger(hook *logrus_sentry.SentryHook) *logrus.Logger {
	logger := logrus.New()
	logger.AddHook(hook)
	return logger
}