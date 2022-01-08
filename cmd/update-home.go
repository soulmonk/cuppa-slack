package main

import (
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"os"
	"soulmonk/cuppa-slack/pkg/drivers"
)

func main() {
	_ = godotenv.Load("./.env")
	app, _, err := drivers.ConnectToSlackViaSocketmode()
	if err != nil {
		panic(err)
	}
	var blocks []slack.Block

	headerText := slack.NewTextBlockObject(slack.MarkdownType, "A simple stack of blocks for the simple sample Block Kit Home tab.", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)
	// Approve and Deny Buttons
	wolBtnTxt := slack.NewTextBlockObject("plain_text", "WOL", false, false)
	wolBtn := slack.NewButtonBlockElement("wol", "commands", wolBtnTxt)

	pingBtnTxt := slack.NewTextBlockObject("plain_text", "Ping", false, false)
	pingBtn := slack.NewButtonBlockElement("ping", "commands", pingBtnTxt)

	actionBlock := slack.NewActionBlock("", wolBtn, pingBtn)

	blocks = append(blocks, headerSection)
	blocks = append(blocks, actionBlock)

	view := slack.HomeTabViewRequest{Type: slack.VTHomeTab, Blocks: slack.Blocks{BlockSet: blocks}}
	if _, err = app.PublishView(os.Getenv("SLACK_USER_ID"), view, ""); err != nil {
		panic(err.Error())
	}
	print("done")
}
