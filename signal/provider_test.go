package signal

import (
    "os"
    "testing"
)

func TestServiceProviderStopSafe(t *testing.T) {
    sp := &ServiceProvider{}
    sp.signalChannel = make(chan os.Signal)
    go func() { <-sp.signalChannel }()
    sp.Stop()
}
