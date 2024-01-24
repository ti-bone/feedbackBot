/*
 * LogMessage.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"feedbackBot/src/config"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"log"
	"os"
)

// LogMessage - logs non-error message to the logs chat
func LogMessage(message string, botInstance *gotgbot.Bot) {
	targetChatId := config.CurrentConfig.LogsID

	if targetChatId == 0 {
		log.SetOutput(os.Stdout)
		log.Println("logs_id is not set in config.json, skipping log event...")
		return
	}

	_, err := botInstance.SendMessage(
		targetChatId,
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
