package signal

import (
	"github.com/goal-web/contracts"
	"os"
	"os/signal"
	"syscall"
)

type ServiceProvider struct {
	signals       []os.Signal
	signalChannel chan os.Signal
	app           contracts.Application
}

func NewService(signals ...os.Signal) contracts.ServiceProvider {
	return &ServiceProvider{signals: signals}
}

func (provider *ServiceProvider) Register(application contracts.Application) {
	provider.app = application
}

func (provider *ServiceProvider) Start() (err error) {
	provider.signalChannel = make(chan os.Signal)
	signal.Notify(provider.signalChannel, provider.signals...)
	for sign := range provider.signalChannel {
		switch sign {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			go func() { provider.app.Stop() }()
		}
	}

	return err
}

func (provider *ServiceProvider) Stop() {
	close(provider.signalChannel)
}
