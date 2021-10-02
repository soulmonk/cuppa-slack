package commands

import (
	"github.com/rs/zerolog/log"
	"os"
)

type wol struct {
	ipAddress  string
	macAddress string
}

func (w wol) Run() {
	log.Debug().
		Str("ipAddress", w.ipAddress).
		Str("macAddress", w.macAddress).
		Msg("Running wol")

	// fail ! forgot about network isolation in kubernetes
	// to be continued
}

func newWol() *wol {
	return &wol{
		ipAddress:  os.Getenv("WOL_IP_ADDRESS"),
		macAddress: os.Getenv("WOL_MAC_ADDRESS"),
	}
}
