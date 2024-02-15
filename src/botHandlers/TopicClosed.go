/*
 * TopicReopened.go
 * Copyright (c) ti-bone 2023-2024
 */

package botHandlers

import (
	"errors"
	"feedbackBot/src/helpers"
	"feedbackBot/src/messages"
	"feedbackBot/users"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func TopicClosed(b *gotgbot.Bot, ctx *ext.Context) error {
	user, err := users.GetUserByTopicId(ctx.EffectiveMessage.MessageThreadId)

	if err != nil && errors.Is(err, messages.UserNotFound) {
		return nil
	} else if err != nil {
		return err
	}

	err = helpers.BanUser(user)

	if err != nil && errors.Is(err, messages.UserAlreadyBanned) {
		return nil
	} else if err != nil {
		return err
	}

	helpers.LogMessage(fmt.Sprintf("#u%d has been banned. Reason: topic with user closed.", user.UserID), b)

	return nil
}
