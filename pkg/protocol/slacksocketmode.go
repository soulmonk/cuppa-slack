package protocol

import (
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"soulmonk/cuppa-slack/pkg/commands"
	"soulmonk/cuppa-slack/pkg/drivers"
)

func SlackSocketMode() {
	api, client, err := drivers.ConnectToSlackViaSocketmode()
	if err != nil {
		log.Fatal().
			Str("error", err.Error()).
			Msg("Unable to connect to slack")
	}

	cmds := commands.Init()

	go func() {
		for evt := range client.Events {
			switch evt.Type {
			case socketmode.EventTypeHello:
				log.Info().Msg("Connected to Slack received hello")
			case socketmode.EventTypeConnecting:
				log.Debug().Msg("Connecting to Slack with Socket Mode...")
			case socketmode.EventTypeConnectionError:
				log.Debug().Msg("Connection failed. Retrying later...")
			case socketmode.EventTypeConnected:
				log.Debug().Msg("Connected to Slack with Socket Mode.")
			case socketmode.EventTypeEventsAPI:
				eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					log.Debug().Interface("evt", evt).Msg("Ignored")
					continue
				}

				log.Debug().Interface("eventsAPIEvent", eventsAPIEvent).Msg("Event received")

				client.Ack(*evt.Request)

				switch eventsAPIEvent.Type {
				case slackevents.CallbackEvent:
					innerEvent := eventsAPIEvent.InnerEvent
					switch ev := innerEvent.Data.(type) {
					case *slackevents.MessageEvent:
						log.Debug().Interface("MessageEvent", ev).Msg("MessageEvent")
						if ev.Text == "Hi" {
							_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
							if err != nil {
								log.Error().Err(err).Interface("type", "MessageEvent").Msg("failed posting message")
							}
						}
					case *slackevents.AppMentionEvent:
						_, _, err := api.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
						if err != nil {
							log.Error().Err(err).Interface("type", "AppMentionEvent").Msg("failed posting message")
						}
					default:
						log.Warn().Interface("MessageEvent", ev).Msg("unsupported innerEvent")
					}
				default:
					log.Warn().Msg("unsupported Events API event received")
					client.Debugf("unsupported Events API event received")
				}
			case socketmode.EventTypeInteractive:
				callback, ok := evt.Data.(slack.InteractionCallback)
				if !ok {
					log.Debug().Interface("evt", evt).Msg("Ignored")
					continue
				}
				log.Debug().Interface("callback", callback).Msg("Interaction received")

				var payload interface{}

				switch callback.Type {
				case slack.InteractionTypeBlockActions:
					// See https://api.slack.com/apis/connections/socket-implement#button

					blockAction := callback.ActionCallback.BlockActions[0]
					if blockAction.Value == "commands" {
						message := cmds.Run(blockAction.ActionID)

						_, err := api.PostEphemeral(callback.User.ID, callback.User.ID, slack.MsgOptionText(message, false))
						if err != nil {
							log.Error().Err(err).Interface("type", "EventTypeSlashCommand").Msg("failed posting ephemeral")
						}
					}

				case slack.InteractionTypeShortcut:
				case slack.InteractionTypeViewSubmission:
					// See https://api.slack.com/apis/connections/socket-implement#modal
				case slack.InteractionTypeDialogSubmission:
				default:

				}

				client.Ack(*evt.Request, payload)
			case socketmode.EventTypeSlashCommand:
				cmd, ok := evt.Data.(slack.SlashCommand)
				if !ok {
					log.Debug().Interface("evt", evt).Msg("Ignored")
					continue
				}

				log.Debug().Interface("cmd", cmd).Msg("Slash command received")

				message := cmds.Run(cmd.Command[1:])
				_, err = api.PostEphemeral(cmd.ChannelID, cmd.UserID, slack.MsgOptionText(message, false))
				if err != nil {
					log.Error().Err(err).Interface("type", "EventTypeSlashCommand").Msg("failed posting ephemeral")
				}

				//var payload interface{}
				client.Ack(*evt.Request)
			default:
				log.Error().
					Str("type", string(evt.Type)).
					Msg("Unexpected event type received")
			}
		}
	}()

	client.Run()
}
