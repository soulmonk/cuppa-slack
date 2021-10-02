package drivers

import (
	"errors"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"log"
	"os"
	"strings"
)

func ConnectToSlackViaSocketmode() (*slack.Client, *socketmode.Client, error) {

	appToken := os.Getenv("SLACK_APP_TOKEN")
	if appToken == "" {
		return nil, nil, errors.New("SLACK_APP_TOKEN must be set")
	}

	if !strings.HasPrefix(appToken, "xapp-") {
		return nil, nil, errors.New("SLACK_APP_TOKEN must have the prefix \"xapp-\"")
	}

	botToken := os.Getenv("SLACK_BOT_TOKEN")
	if botToken == "" {
		return nil, nil, errors.New("SLACK_BOT_TOKEN must be set")
	}

	if !strings.HasPrefix(botToken, "xoxb-") {
		return nil, nil, errors.New("SLACK_BOT_TOKEN must have the prefix \"xoxb-\"")
	}

	debug := os.Getenv("SLACK_SDK_DEBUG")
	isDebug := false
	if debug == "true" {
		isDebug = true
	}

	// LoggerApi := zerolog.New(os.Stderr).With().Str("process", "api").Timestamp().Logger()

	api := slack.New(
		botToken,
		slack.OptionDebug(isDebug),
		slack.OptionAppLevelToken(appToken),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(isDebug),
		socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
	)

	return api, client, nil
}
