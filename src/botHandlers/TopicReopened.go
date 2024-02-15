/*
 * TopicReopened.go
 * Copyright (c) ti-bone 2023-2024
 */

package botHandlers

import (
	"errors"
	"feedbackBot/src/constants"
	"feedbackBot/src/helpers"
	"feedbackBot/src/users"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func TopicReopened(b *gotgbot.Bot, ctx *ext.Context) error {
	user, err := users.GetUserByTopicId(ctx.EffectiveMessage.MessageThreadId)

	if err != nil && errors.Is(err, constants.UserNotFound) {
		return nil
	} else if err != nil {
		return err
	}

	err = helpers.UnbanUser(user)

	if err != nil && errors.Is(err, constants.UserNotBanned) {
		return nil
	} else if err != nil {
		return err
	}

	helpers.LogMessage(fmt.Sprintf("Ban lifted from #u%d. Reason: topic with user reopened.", user.UserID), b)

	return nil
}
