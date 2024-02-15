/*
 * Start.go
 * Copyright (c) ti-bone 2023-2024
 */

package commands

import (
	"feedbackBot/src/config"
	"feedbackBot/src/rates"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	var err error

	// Check if user is not rate-limited and welcome message is enabled
	if rates.Check(ctx.EffectiveChat.Id, 10) && config.CurrentConfig.Welcome.Enabled {
		// Send welcome message
		_, err = ctx.EffectiveMessage.Reply(
			b,
			config.CurrentConfig.Welcome.Message,
			&gotgbot.SendMessageOpts{ParseMode: "HTML"},
		)
	}

	return err
}
