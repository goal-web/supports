package logs

import (
	"github.com/apex/log"
	"github.com/goal-web/contracts"
)

func WithFields(fields contracts.Fields) contracts.Logger {
	return &ApexLogger{Entry: log.WithFields(log.Fields(fields))}
}

func WithError(err error) contracts.Logger {
	return &ApexLogger{Entry: log.WithError(err)}
}

func WithException(exception contracts.Exception) contracts.Logger {
	return &ApexLogger{Entry: log.WithError(exception)}
}

func Default() contracts.Logger {
	return &ApexLogger{Entry: log.WithFields(log.Fields(contracts.Fields{}))}
}

func WithInterface(value any) contracts.Logger {
	return WithField("value", value)
}

func WithField(key string, value any) contracts.Logger {
	return &ApexLogger{Entry: log.WithField(key, value)}
}
