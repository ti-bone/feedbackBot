/*
 * Response.go
 * Copyright (c) ti-bone 2023
 */

package botHandlers

import (
	"feedbackBot/src/db"
	"feedbackBot/src/helpers"
	"feedbackBot/src/models"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Response(b *gotgbot.Bot, ctx *ext.Context) error {
	var senderUser models.User

	db.Connection.Where("user_id = ?", ctx.EffectiveSender.Id()).First(&senderUser)

	if !senderUser.IsAdmin {
		return nil
	}

	var user models.User

	db.Connection.Where("topic_id = ?", ctx.EffectiveMessage.MessageThreadId).First(&user)

	if user.TopicID != 0 && !user.IsBanned {
		_, err := b.CopyMessage(user.UserID, ctx.EffectiveChat.Id, ctx.EffectiveMessage.MessageId, &gotgbot.CopyMessageOpts{})

		if err != nil {
			secondErr := helpers.LogError(err.Error(), b, ctx)

			if secondErr != nil {
				return secondErr
			}

			return err
		}
	}

	return nil
}
