/*
 * Response.go
 * Copyright (c) ti-bone 2023-2024
 */

package handlers

import (
	"errors"
	"feedbackBot/src/constants"
	"feedbackBot/src/db"
	"feedbackBot/src/messages"
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
		id, err := b.CopyMessage(
			user.UserID,
			ctx.EffectiveChat.Id,
			ctx.EffectiveMessage.MessageId,
			&gotgbot.CopyMessageOpts{ProtectContent: user.IsProtected},
		)

		var tgErr *gotgbot.TelegramError

		if errors.As(err, &tgErr) {
			// If thread not found - try to recreate topic
			if tgErr.Description == "Forbidden: bot was blocked by the user" {
				_, err = ctx.EffectiveMessage.Reply(b, constants.BotUserBlocked.Error(), &gotgbot.SendMessageOpts{})
				return err
			}
		}

		if err != nil {
			return err
		}

		// Save the message identifiers relation
		return messages.StoreMessage(
			models.Message{
				UserID:           user.UserID,
				UserMessageId:    id.MessageId,
				SupportMessageId: ctx.EffectiveMessage.MessageId,
				SupportChatId:    ctx.EffectiveChat.Id,
				IsOutgoing:       true,
			},
		)
	}

	return err
}
