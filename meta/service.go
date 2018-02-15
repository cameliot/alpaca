package meta

import (
	"github.com/cameliot/alpaca"

	"log"
	"time"
)

type MetaSvc struct {
	iama      IamaPayload
	startedAt time.Time
}

func NewMetaSvc(
	handle string,
	name string,
	version string,
	description string,
) *MetaSvc {
	svc := &MetaSvc{
		iama: IamaPayload{
			Handle:      handle,
			Name:        name,
			Version:     version,
			Description: description,
			StartedAt:   time.Now().UTC().UnixNano() / 1000000,
		},
	}

	return svc
}

func (self *MetaSvc) Handle(actions alpaca.Actions, dispatch alpaca.Dispatch) {
	log.Println("Started  _meta actions handling")

	for action := range actions {
		switch action.Type {
		case PING:
			self.handlePing(action, dispatch)
			break
		case WHOIS:
			self.handleWhois(action, dispatch)
			break
		}
	}
}

/*
 Handle PING,
 Reply only if the requested service is a wildcard ("*") or
 identified by the service handler
*/
func (self *MetaSvc) handlePing(
	action alpaca.Action,
	dispatch alpaca.Dispatch,
) {
	payload := ""
	action.DecodePayload(&payload)

	// Are we pinged?
	if payload != "*" && payload != self.iama.Handle {
		return
	}

	// Reply with PONG
	dispatch(Pong(self.iama.Handle))
}

/*
 Handle WHOIS,

 Reply only if the requested service is a wildcard ("*") or
 identified by the service handler

 Provide information about this service
*/
func (self *MetaSvc) handleWhois(
	action alpaca.Action,
	dispatch alpaca.Dispatch,
) {
	payload := ""
	action.DecodePayload(&payload)

	// Are we pinged?
	if payload != "*" && payload != self.iama.Handle {
		return
	}

	// Reply with IAMA
	dispatch(Iama(self.iama))
}
