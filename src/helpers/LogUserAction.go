/*
 * LogUserAction.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"feedbackBot/src/config"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"html"
	"math"
)

func LogUserAction(message string, b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveUser
	chat := ctx.EffectiveChat

	logsID := config.CurrentConfig.LogsID

	logMessage := fmt.Sprintf(
		"%s\nUserID: <code>%d</code>\nFirst Name: <code>%s</code>",
		message,
		user.Id,
		html.EscapeString(user.FirstName))

	if user.LastName != "" {
		logMessage = fmt.Sprintf("%s\nLast Name: <code>%s</code>", logMessage, html.EscapeString(user.LastName))
	}

	if user.Username != "" {
		logMessage = fmt.Sprintf("%s\nUsername: @%s", logMessage, html.EscapeString(user.Username))
	}

	if user.LanguageCode != "" {
		logMessage = fmt.Sprintf(
			"%s\nLanguage: <code>%s</code>",
			logMessage,
			html.EscapeString(user.LanguageCode),
		)
	}

	if chat.Type != "private" {
		logMessage = fmt.Sprintf(
			"%s\nChat Type: <code>%s</code>\nChat Name: <code>%s</code>\nChat ID: <code>%d</code>\n#ch%.f",
			logMessage,
			html.EscapeString(chat.Type),
			html.EscapeString(chat.Title),
			chat.Id,
			math.Abs(float64(chat.Id)),
		)
	}

	logMessage = fmt.Sprintf(
		"%s\n#u%d",
		logMessage,
		user.Id,
	)

	_, SendToLogsErr := b.SendMessage(
		logsID,
		logMessage,
		&gotgbot.SendMessageOpts{
			ParseMode:       "html",
			MessageThreadId: config.CurrentConfig.LogsTopicID,
		},
	)

	if SendToLogsErr != nil {
		return SendToLogsErr
	}

	return nil
}
