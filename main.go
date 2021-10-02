package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"soulmonk/cuppa-slack/pkg/protocol"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	_ = godotenv.Load("./.env")

	go protocol.Rest()
	protocol.SlackSocketMode()
}
