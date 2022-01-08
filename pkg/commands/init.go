package commands

import "github.com/rs/zerolog/log"

type Commands struct {
	wol  *wol
	ping *ping
}

func (c *Commands) Run(name string) (message string) {
	switch name {
	case "wol":
		message = "Done"
		if err := c.wol.run(); err != nil {
			log.Error().Err(err).Interface("from", "Commands").Interface("name", name).Msg("cannot run command")
			message = "Error, check logs"
		}
	case "ping":
		message = "Live"
		if ok, err := c.ping.run(); err != nil {
			log.Error().Err(err).Interface("from", "Commands").Interface("name", name).Msg("cannot run command")
			message = "Error, check logs"
		} else if !ok {
			message = "Down"
		}

	default:
		message = "Unknown command!"
	}
	return
}

func Init() Commands {
	// todo
	// to map
	// {'/commandName': commandInstance()}
	return Commands{
		newWol(),
		newPing(),
	}
}
