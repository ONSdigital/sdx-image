// Package log wraps "zerolog" with convenience calls for logging in SDX.
package log

import (
	"github.com/rs/zerolog/log"
)

func Init() {
	log.Logger = log.With().Str("app", "sdx-image").Logger()
}

func Info(msg string, txId ...string) {
	event := log.Log().Str("message", msg).Str("severity", "INFO")
	if txId != nil && len(txId) != 0 {
		event.Str("tx_id", txId[0])
	}
	event.Send()
}

func Error(msg string, err error, txId ...string) {
	event := log.Log().Str("message", msg).Str("severity", "ERROR").Str("error", err.Error())
	if txId != nil && len(txId) != 0 {
		event.Str("tx_id", txId[0])
	}
	event.Send()
}
