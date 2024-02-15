/*
 * IsResolvable.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"feedbackBot/src/config"
	"feedbackBot/src/constants"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func IsResolvable(ctx *ext.Context, b *gotgbot.Bot) (bool, bool, error) {
	args := ctx.Args()
	topicId := ctx.EffectiveMessage.MessageThreadId
	isValidTopicMessage := topicId != 0 && topicId != config.CurrentConfig.LogsTopicID

	if !isValidTopicMessage && len(args) <= 1 {
		_, err := ctx.EffectiveMessage.Reply(
			b,
			constants.UserNotSpecified.Error()+"\nHint: you can send this command to any user-related topic.",
			&gotgbot.SendMessageOpts{},
		)

		return false, false, err
	}

	return true, isValidTopicMessage, nil
}
