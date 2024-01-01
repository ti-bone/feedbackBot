/*
 * Ban.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"feedbackBot/src/db"
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

	if user.IsBanned {
		_, err = ctx.EffectiveMessage.Reply(b, messages.UserAlreadyBanned.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	user.IsBanned = true
	db.Connection.Updates(&user)

	_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("#u%d has been banned.", user.UserID), &gotgbot.SendMessageOpts{})
	return err
}
