/*
 * Reaction.go
 * Copyright (c) ti-bone 2023-2024
 */

package handlers

import (
	"errors"
	"feedbackBot/src/config"
	"feedbackBot/src/constants"
	"feedbackBot/src/helpers"
	"feedbackBot/src/messages"
	"feedbackBot/src/reactions"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Reaction(b *gotgbot.Bot, ctx *ext.Context) error {
	reaction := ctx.MessageReaction

	if reaction.Chat.Id == config.CurrentConfig.LogsID {
		// If the reaction was set in the logs chat, process it as outgoing reaction to the user
		message, err := messages.GetMessageBySupportId(reaction.MessageId, reaction.Chat.Id)

		if errors.Is(err, constants.MessageNotFound) {
			return nil
		}

		if err != nil {
			return helpers.LogError(err.Error(), b, ctx)
		}

		return reactions.ProcessUpdateReactions(reaction, b, message.UserId, message.UserMessageId)
	} else if reaction.User != nil && reaction.Chat.Id == reaction.User.Id {
		// If the reaction was set by the user, process it as incoming reaction to the support
		message, err := messages.GetMessageByUserId(reaction.MessageId, reaction.User.Id)

		if errors.Is(err, constants.MessageNotFound) {
			return nil
		}

		if err != nil {
			return helpers.LogError(err.Error(), b, ctx)
		}

		return reactions.ProcessUpdateReactions(reaction, b, message.SupportChatId, message.SupportMessageId)
	}

	// Ignore others
	return nil
}
