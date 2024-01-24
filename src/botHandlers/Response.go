/*
 * Response.go
 * Copyright (c) ti-bone 2023-2024
 */

package botHandlers

import (
	"feedbackBot/src/db"
	"feedbackBot/src/models"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

func Response(b *gotgbot.Bot, ctx *ext.Context) error {
	// Check if the message is a service message, describing topic action
	if message.TopicAction(ctx.EffectiveMessage) {
		return nil
	}

	var err error
	var user models.User

	// Get target user from the DB by topic ID
	db.Connection.Where("topic_id = ?", ctx.EffectiveMessage.MessageThreadId).First(&user)

	// If user is not found, return
	if user.TopicID != 0 && !user.IsBanned {
		_, err = b.CopyMessage(
			user.UserID,
			ctx.EffectiveChat.Id,
			ctx.EffectiveMessage.MessageId,
			&gotgbot.CopyMessageOpts{ProtectContent: user.IsProtected},
		)
	}

	return err
}
