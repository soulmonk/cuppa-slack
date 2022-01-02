package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"os"
	"soulmonk/cuppa-slack/pkg/protocol"
)

func main() {
	strLogLevel := os.Getenv("LOG_LEVEL")
	if strLogLevel == "" {
		strLogLevel = "debug"
	}
	logLevel, err := zerolog.ParseLevel(strLogLevel)
	if err != nil {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	_ = godotenv.Load("./.env")

	//go protocol.Rest()
	protocol.SlackSocketMode()
}
