/*
 * Message.go
 * Copyright (c) ti-bone 2023-2024
 */

package handlers

import (
	"errors"
	"feedbackBot/src/config"
	"feedbackBot/src/db"
	"feedbackBot/src/helpers"
	"feedbackBot/src/messages"
	"feedbackBot/src/models"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"html"
	"log"
	"os"
	"time"
)

func Message(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveSender.Id() == b.Id {
		return nil
	}

	var handleMessage func(int) error

	handleMessage = func(depth int) error {
		if depth > 5 {
			// Return an error when reaching the maximum depth
			return errors.New("maximum recursion depth reached")
		}

		var user models.User

		res := db.Connection.Where("user_id = ?", ctx.EffectiveUser.Id).First(&user)

		if res.Error != nil {
			// Log the error to admin topic
			err := helpers.LogError(
				fmt.Sprintf("Got DB error: %v, retrying(attempt %d)...", res.Error, depth),
				b, ctx,
			)

			if err != nil {
				log.SetOutput(os.Stderr)
				log.Printf("failed to log error: %v", err.Error())
			}

			// Wait for 2s and retry
			time.Sleep(2 * time.Second)

			// Retry
			return handleMessage(depth + 1)
		}

		if user.IsBanned {
			return nil
		}

		if user.TopicID == 0 {
			return handleNoTopic(b, ctx, &user, handleMessage, depth)
		}

		// Forward message to the user's topic
		response, err := b.ForwardMessage(
			config.CurrentConfig.LogsID,
			ctx.EffectiveChat.Id,
			ctx.EffectiveMessage.MessageId,
			&gotgbot.ForwardMessageOpts{
				MessageThreadId: user.TopicID,
			},
		)

		message := &models.Message{
			UserID:        user.UserID,
			UserMessageId: ctx.EffectiveMessage.MessageId,
			IsOutgoing:    false,
		}

		// Call to response may produce panic, because response could be nil, so we check it
		if response != nil {
			message.SupportMessageId = response.MessageId
		}

		var tgErr *gotgbot.TelegramError

		if errors.As(err, &tgErr) {
			// If thread not found - try to recreate topic
			if tgErr.Description == "Bad Request: message thread not found" {
				return handleNoTopic(b, ctx, &user, handleMessage, depth)
			}
		}

		// If failed, try to copy message
		// (can be useful if the user has SCAM flag, Telegram doesn't allow to forward messages from such users
		if err != nil {
			id, err := b.CopyMessage(
				config.CurrentConfig.LogsID,
				ctx.EffectiveChat.Id,
				ctx.EffectiveMessage.MessageId,
				&gotgbot.CopyMessageOpts{
					MessageThreadId: user.TopicID,
				},
			)

			message.SupportMessageId = id.MessageId

			if err != nil {
				return err
			}
		}

		// Store message info in DB
		return messages.StoreMessage(*message)
	}

	return handleMessage(1)
}

func handleNoTopic(
	b *gotgbot.Bot, ctx *ext.Context,
	user *models.User, handleMessage func(int) error,
	currentDepth int,
) error {
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
		return err
	}

	var username string

	if ctx.EffectiveSender.User.Username != "" {
		username = "\nUsername: @" + ctx.EffectiveSender.User.Username
	}

	_, err = b.SendMessage(
		config.CurrentConfig.LogsID,
		fmt.Sprintf(
			"This topic with ID <code>%d</code> belongs to user <code>%s</code> %sID: <code>%d</code>%s",
			topic.MessageThreadId,
			html.EscapeString(ctx.EffectiveUser.FirstName),
			"<code>"+html.EscapeString(ctx.EffectiveUser.LastName)+"</code> ",
			ctx.EffectiveUser.Id,
			username,
		),
		&gotgbot.SendMessageOpts{
			ParseMode:       "HTML",
			MessageThreadId: topic.MessageThreadId,
		},
	)

	if err != nil {
		// Delete topic(no need for it, because first message failed to send)
		_, _ = b.DeleteForumTopic(config.CurrentConfig.LogsID, topic.MessageThreadId, &gotgbot.DeleteForumTopicOpts{})

		return err
	}

	// Set the topic ID to the user and write it to the DB
	user.TopicID = topic.MessageThreadId
	db.Connection.Where("user_id = ?", user.UserID).Updates(&user)

	return handleMessage(currentDepth + 1)
}
