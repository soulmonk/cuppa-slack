package commands

import (
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"strings"
)

type wol struct {
	ipAddress  string
	macAddress string
}

func (w wol) Run() error {
	log.Debug().
		Str("ipAddress", w.ipAddress).
		Str("macAddress", w.macAddress).
		Msg("Running wol")

	out, _ := exec.Command("ping", w.ipAddress, "-c 5", "-i 3", "-w 10").Output()
	isOnline := strings.Contains(string(out), "Destination Host Unreachable")
	if isOnline {
		log.Debug().Msg("TANGO DOWN")
	} else {
		log.Debug().Msg("IT'S ALIVEEE")
	}

	//cmd := exec.Command("wakeonlan", "-i", w.ipAddress, w.macAddress)
	// for now use external app
	cmd := exec.Command("wakeonlan", w.macAddress)
	stdout, err := cmd.Output()

	if err != nil {
		log.Error().
			Interface("error", err.Error()).
			Msg("Error running wakeonlan")
		return err
	}
	log.Info().
		Str("output", string(stdout)).
		Msg("Result wol")
	return nil
}

func newWol() *wol {
	return &wol{
		ipAddress:  os.Getenv("WOL_IP_ADDRESS"),
		macAddress: os.Getenv("WOL_MAC_ADDRESS"),
	}
}
