/*
 * Unban.go
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

func Unban(b *gotgbot.Bot, ctx *ext.Context) error {
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

	if !user.IsBanned {
		_, err = ctx.EffectiveMessage.Reply(b, messages.UserNotBanned.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	db.Connection.Model(&user).Update("is_banned", false)

	_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Ban lifted from #u%d.", user.UserID), &gotgbot.SendMessageOpts{})
	return err
}
