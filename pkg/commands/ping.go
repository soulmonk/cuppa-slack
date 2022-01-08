package commands

import (
	"github.com/rs/zerolog/log"
	"os"
	"os/exec"
	"strings"
)

type ping struct {
	ipAddr string
}

func (p *ping) run() (bool, error) {
	log.Debug().
		Str("ipAddr", p.ipAddr).
		Msg("Running ping")
	out, err := exec.Command("ping", p.ipAddr, "-c 3", "-i 3").Output()
	if err != nil {
		return false, err
	}
	isOnline := !strings.Contains(string(out), "Destination Host Unreachable")
	return isOnline, nil
}

func newPing() *ping {
	return &ping{
		ipAddr: os.Getenv("WOL_IP_ADDRESS"), // command parameter, no need for this specific app
	}
}
