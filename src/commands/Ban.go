/*
 * Ban.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"errors"
	"feedbackBot/src/helpers"
	"feedbackBot/src/messages"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Ban(b *gotgbot.Bot, ctx *ext.Context) error {
	args := ctx.Args()

	if len(args) <= 1 {
		_, err := ctx.EffectiveMessage.Reply(b, messages.UserNotSpecified.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	user, err := helpers.ParseInputUser(args[1])

	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	err = helpers.BanUser(user)

	if err != nil && errors.Is(err, messages.UserAlreadyBanned) {
		_, err = ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
	} else if err != nil {
		return err
	}

	_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("#u%d has been banned.", user.UserID), &gotgbot.SendMessageOpts{})
	return err
}
