package conductor

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jpoz/switchboard/pkg/config"
	"github.com/jpoz/switchboard/pkg/proxy"
	log "github.com/sirupsen/logrus"
)

// Conductor is the main object incharge of:
//  - Starting proxies
//  - Keeping state of those proxies
//  - Stopping proxies when SIG command received
type Conductor struct {
	Config config.Config
}

type namedError struct {
	name string
	err  error
}

// New should be used to create any Conductors
func New(c config.Config) *Conductor {
	return &Conductor{
		Config: c,
	}
}

// Start will attempt to startup all proxies and wait for sigint
func (c Conductor) Start() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println() // this is to clear the line after the ^C
		log.Infof("Received: %s", sig)
		done <- true
	}()

	errsig := make(chan namedError)

	log.Info("Starting conductor 🚊")
	if len(c.Config.LocalProxies) > 0 {
		for name, lc := range c.Config.LocalProxies {
			esig := make(chan error)
			lp, err := proxy.NewLocal(name, lc.ProxyAddr, lc.ConnectingAddr, esig)
			if err != nil {
				log.Errorf("%s failed to start: %s", name, err)
				return
			}

			go func(name string) {
				go lp.Listen()
				err := <-esig
				errsig <- namedError{
					name: name,
					err:  err,
				}
			}(name)
		}
	}

	select {
	case <-done:
		log.Info("Conductor exiting 👋")
	case namedErr := <-errsig:
		log.Errorf("Error from %s: %s", namedErr.name, namedErr.err)
	}
}
