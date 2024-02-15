/*
 * SendLargeText.go
 * Copyright (c) ti-bone 2023-2024
 */

package helpers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func SendLargeText(
	ctx *ext.Context, b *gotgbot.Bot,
	text string, options *gotgbot.SendMessageOpts,
) error {
	if len(text) > 2048 {
		for len(text) > 2048 {
			_, err := ctx.EffectiveMessage.Reply(
				b,
				text[:2048],
				options,
			)

			if err != nil {
				return err
			}

			text = text[2048:]
		}
	}

	_, err := ctx.EffectiveMessage.Reply(
		b,
		text,
		options,
	)

	return err
}
