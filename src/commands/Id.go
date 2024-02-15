/*
 * Id.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"feedbackBot/src/rates"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Id(b *gotgbot.Bot, ctx *ext.Context) error {
	var err error
	// Check if chat is not rate-limited
	if rates.Check(ctx.EffectiveChat.Id, 60) {
		// Send message with ID
		_, err = ctx.EffectiveMessage.Reply(
			b,
			fmt.Sprintf("Chat ID: <code>%d</code>", ctx.EffectiveChat.Id),
			&gotgbot.SendMessageOpts{ParseMode: "HTML"},
		)
	}

	return err
}
