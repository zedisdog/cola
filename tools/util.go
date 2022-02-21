package tools

import (
	"os"
	"os/signal"
)

func Wait(closeFunc ...func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		for _, cls := range closeFunc {
			cls()
		}
	}
}
