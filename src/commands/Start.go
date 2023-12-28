/*
 * Start.go
 * Copyright (c) ti-bone 2023
 */

package commands

import (
	"feedbackBot/src/config"
	"feedbackBot/src/rateLimiter"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	if rateLimiter.Check(ctx.EffectiveChat.Id) && config.CurrentConfig.Welcome.Enabled {
		_, err := ctx.EffectiveMessage.Reply(b, config.CurrentConfig.Welcome.Message, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
		return err
	}

	return nil
}
