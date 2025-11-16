package logs

import "testing"

func TestApexLoggerDebugSwitch(t *testing.T) {
    l := Default()
    Debug = false
    l.Debug("skip")
    Debug = true
    l.Debug("emit")
    l.Info("info")
    l.Error("error")
}