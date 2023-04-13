package logs

import (
	"github.com/apex/log"
	"github.com/goal-web/contracts"
)

var Debug = false

type ApexLogger struct {
	Entry *log.Entry
}

func (apex *ApexLogger) WithFields(m contracts.Fields) contracts.Logger {
	if apex == nil || apex.Entry == nil {
		return &ApexLogger{
			Entry: log.WithFields(log.Fields(m)),
		}
	}

	apex.Entry = apex.Entry.WithFields(log.Fields(m))

	return apex
}

func (apex *ApexLogger) WithField(key string, value any) contracts.Logger {
	if apex == nil || apex.Entry == nil {
		return &ApexLogger{
			Entry: log.WithField(key, value),
		}
	}

	apex.Entry = apex.Entry.WithField(key, value)

	return apex
}

func (apex *ApexLogger) WithError(err error) contracts.Logger {
	if apex == nil || apex.Entry == nil {
		return &ApexLogger{
			Entry: log.WithError(err),
		}
	}

	apex.Entry = apex.Entry.WithError(err)

	return apex
}

func (apex *ApexLogger) WithException(err contracts.Exception) contracts.Logger {
	if apex == nil || apex.Entry == nil {
		return &ApexLogger{
			Entry: log.WithError(err),
		}
	}

	apex.Entry = apex.Entry.WithError(err)

	return apex
}

func (apex *ApexLogger) Info(msg string) {
	apex.Entry.Info(msg)
}

func (apex *ApexLogger) Warn(msg string) {
	apex.Entry.Warn(msg)
}

func (apex *ApexLogger) Debug(msg string) {
	if Debug {
		apex.Entry.Debug(msg)
	}
}

func (apex *ApexLogger) Error(msg string) {
	apex.Entry.Error(msg)
}

func (apex *ApexLogger) Fatal(msg string) {
	apex.Entry.Fatal(msg)
}
