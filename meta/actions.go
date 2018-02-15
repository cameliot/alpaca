package meta

import (
	"github.com/cameliot/alpaca"
	"time"
)

const PING = "@meta/PING"
const PONG = "@meta/PONG"
const WHOIS = "@meta/WHOIS"
const IAMA = "@meta/IAMA"

type PongPayload struct {
	TimestampMs int64  `json:"timestamp"`
	Handle      string `json:"handle"`
}

/*
Decode millisecond timestamp into time.Time
*/
func decodeTimestampMs(t int64) time.Time {
	sec := t / 1000
	nsec := 1000000 * (t % 1000)

	return time.Unix(sec, nsec).UTC()
}

/*
Encode millisecond (UTC) timestamp
*/
func encodeTimestampMs(t time.Time) int64 {
	return t.UTC().UnixNano() / 1000000
}

/*
Decode int64 millisecond timestamp
*/
func (payload PongPayload) Timestamp() time.Time {
	return decodeTimestampMs(payload.TimestampMs)
}

func DecodePong(action alpaca.Action) PongPayload {
	payload := PongPayload{}
	action.DecodePayload(&payload)

	return payload
}

type IamaPayload struct {
	Name        string `json:"name"`
	Handle      string `json:"handle"`
	Version     string `json:"version"`
	Description string `json:"description"`
	StartedAtMs int64  `json:"started_at"`
}

func (payload IamaPayload) StartedAt() time.Time {
	return decodeTimestampMs(payload.StartedAtMs)
}

func DecodeIama(action alpaca.Action) IamaPayload {
	payload := IamaPayload{}
	action.DecodePayload(&payload)

	return payload
}

// Action Creators

func Ping(handle string) alpaca.Action {
	return alpaca.Action{
		Type:    PING,
		Payload: handle,
	}
}

func Pong(handle string) alpaca.Action {
	return alpaca.Action{
		Type: PONG,
		Payload: PongPayload{
			Handle:      handle,
			TimestampMs: encodeTimestampMs(time.Now()),
		},
	}
}

func Whois(handle string) alpaca.Action {
	return alpaca.Action{
		Type:    WHOIS,
		Payload: handle,
	}
}

func Iama(manifest IamaPayload) alpaca.Action {
	return alpaca.Action{
		Type:    IAMA,
		Payload: manifest,
	}
}
