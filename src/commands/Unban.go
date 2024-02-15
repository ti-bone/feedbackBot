/*
 * Unban.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"errors"
	"feedbackBot/src/constants"
	"feedbackBot/src/helpers"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Unban(b *gotgbot.Bot, ctx *ext.Context) error {
	user, err := helpers.ResolveUser(ctx, b)

	if err != nil || user == nil {
		return err
	}

	err = helpers.UnbanUser(user)

	if err != nil && errors.Is(err, constants.UserNotBanned) {
		_, err = ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
	} else if err != nil {
		return err
	}

	_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Ban lifted from #u%d.", user.UserId), &gotgbot.SendMessageOpts{})
	return err
}
