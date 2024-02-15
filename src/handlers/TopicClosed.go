/*
 * TopicReopened.go
 * Copyright (c) ti-bone 2023-2024
 */

package handlers

import (
	"errors"
	"feedbackBot/src/constants"
	"feedbackBot/src/helpers"
	"feedbackBot/src/users"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func TopicClosed(b *gotgbot.Bot, ctx *ext.Context) error {
	user, err := users.GetUserByTopicId(ctx.EffectiveMessage.MessageThreadId)

	if err != nil && errors.Is(err, constants.UserNotFound) {
		return nil
	} else if err != nil {
		return err
	}

	err = helpers.BanUser(user)

	if err != nil && errors.Is(err, constants.UserAlreadyBanned) {
		return nil
	} else if err != nil {
		return err
	}

	helpers.LogMessage(fmt.Sprintf("#u%d has been banned. Reason: topic with user closed.", user.UserId), b)

	return nil
}
