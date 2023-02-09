package log

import (
	"github.com/rs/zerolog/log"
)

func Init() {
	log.Logger = log.With().Str("app", "sdx-image").Logger()
}

func Info(msg string, txId ...string) {
	if txId != nil && len(txId) != 0 {
		log.Info().Str("event", msg).Str("tx_id", txId[0]).Send()
	} else {
		log.Info().Str("event", msg).Send()
	}
}

func Error(msg string, err error, txId ...string) {
	if txId != nil && len(txId) != 0 {
		log.Error().Str("event", msg).Str("error", err.Error()).Str("tx_id", txId[0]).Send()
	} else {
		log.Error().Str("event", msg).Str("error", err.Error()).Send()
	}
}
