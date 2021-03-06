package hetzner

import (
	"log"
	"time"
)

type Alerter struct {
	Send    func(Server) error
	Realert time.Duration
	alerted map[int]time.Time
}

func (a *Alerter) Start() chan Server {
	a.alerted = map[int]time.Time{}
	srvchan := make(chan Server)

	go func() {
		for srv := range srvchan {
			if time.Since(a.alerted[srv.ID]) < a.Realert {
				continue
			}

			a.alerted[srv.ID] = time.Now()
			err := a.Send(srv)
			if err != nil {
				log.Printf("Error sending alert: %v", err)
			}
		}
	}()

	return srvchan
}
