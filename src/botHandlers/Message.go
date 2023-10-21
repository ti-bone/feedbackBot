/*
 * Message.go
 * Copyright (c) ti-bone 2023
 */

package botHandlers

import (
	"feedbackBot/src/config"
	"feedbackBot/src/db"
	"feedbackBot/src/helpers"
	"feedbackBot/src/models"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"html"
)

func Message(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveSender.Id() == b.Id || ctx.EffectiveMessage.Text == "/start" {
		return nil
	}

	var user models.User

	res := db.Connection.Where("user_id = ?", ctx.EffectiveUser.Id).First(&user)

	if res.Error != nil {
		err := helpers.LogError(res.Error.Error(), b, ctx)

		if err != nil {
			return err
		}

		return res.Error
	}

	if user.IsBanned {
		return nil
	}

	if user.TopicID == 0 {
		topic, err := b.CreateForumTopic(
			config.CurrentConfig.LogsID,
			fmt.Sprintf(
				"%s [%d]",
				ctx.EffectiveUser.FirstName,
				ctx.EffectiveUser.Id,
			),
			&gotgbot.CreateForumTopicOpts{},
		)

		if err != nil {
			secondErr := helpers.LogError(err.Error(), b, ctx)

			if secondErr != nil {
				return secondErr
			}

			return err
		}

		_, err = b.SendMessage(
			config.CurrentConfig.LogsID,
			fmt.Sprintf(
				"This topic with ID <code>%d</code> belongs to user <code>%s</code> %sID: <code>%d</code>",
				topic.MessageThreadId,
				html.EscapeString(ctx.EffectiveUser.FirstName),
				"<code>"+html.EscapeString(ctx.EffectiveUser.LastName)+"</code> ",
				ctx.EffectiveUser.Id,
			),
			&gotgbot.SendMessageOpts{
				ParseMode:       "HTML",
				MessageThreadId: topic.MessageThreadId,
			},
		)

		if err != nil {
			secondErr := helpers.LogError(err.Error(), b, ctx)

			_, _ = b.DeleteForumTopic(config.CurrentConfig.LogsID, topic.MessageThreadId, &gotgbot.DeleteForumTopicOpts{})

			if secondErr != nil {
				return secondErr
			}

			return err
		}

		user.TopicID = topic.MessageThreadId

		db.Connection.Where("user_id = ?", user.UserID).Updates(&user)
	}

	_, err := b.ForwardMessage(
		config.CurrentConfig.LogsID,
		ctx.EffectiveChat.Id,
		ctx.EffectiveMessage.MessageId,
		&gotgbot.ForwardMessageOpts{
			MessageThreadId: user.TopicID,
		},
	)

	if err != nil {
		// Try to copyMessage
		_, err := b.CopyMessage(
			config.CurrentConfig.LogsID,
			ctx.EffectiveChat.Id,
			ctx.EffectiveMessage.MessageId,
			&gotgbot.CopyMessageOpts{
				MessageThreadId: user.TopicID,
			},
		)

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
