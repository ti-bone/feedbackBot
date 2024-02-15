/*
 * DelMessage.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"errors"
	"feedbackBot/src/constants"
	"feedbackBot/src/messages"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
	"os"
)

func DelMessage(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveMessage.ReplyToMessage == nil {
		_, err := ctx.EffectiveMessage.Reply(b, constants.NoMessageToDelete.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	replyTo := ctx.EffectiveMessage.ReplyToMessage

	if replyTo.MessageId == replyTo.MessageThreadId {
		_, err := ctx.EffectiveMessage.Reply(b, constants.NoMessageToDelete.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	message, err := messages.GetMessageBySupportId(replyTo.MessageId, ctx.EffectiveChat.Id)

	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	_, err = b.DeleteMessage(message.UserID, message.UserMessageId, &gotgbot.DeleteMessageOpts{})

	// If delete failed due to message not found - output error message and return
	var tgErr *gotgbot.TelegramError

	if errors.As(err, &tgErr) {
		if tgErr.Description == "Bad Request: message to delete not found" {
			_, err = ctx.EffectiveMessage.Reply(b, constants.MessageAlreadyDeleted.Error(), &gotgbot.SendMessageOpts{})
			return err
		}
	}

	if err != nil {
		log.SetOutput(os.Stderr)
		log.Println(err.Error())

		_, err := ctx.EffectiveMessage.Reply(b, constants.InternalError.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	_, err = ctx.EffectiveMessage.Reply(
		b,
		fmt.Sprintf(
			constants.MessageDeleted,
			message.UserMessageId, message.UserID,
		),
		&gotgbot.SendMessageOpts{},
	)

	return err
}
