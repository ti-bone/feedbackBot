/*
 * LogMessage.go
 * Copyright (c) ti-bone 2023
 */

package helpers

import (
	"feedbackBot/src/config"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"log"
	"os"
)

func LogMessage(message string, botInstance *gotgbot.Bot) {
	_, err := botInstance.SendMessage(
		config.CurrentConfig.LogsID,
		message,
		&gotgbot.SendMessageOpts{
			MessageThreadId: config.CurrentConfig.LogsTopicID,
		},
	)

	if err != nil {
		log.SetOutput(os.Stderr)
		log.Println(err)
	}
}
