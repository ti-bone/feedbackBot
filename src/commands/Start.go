package commands

import (
	"feedbackBot/src/config"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	if config.CurrentConfig.Welcome.Enabled {
		_, err := ctx.EffectiveMessage.Reply(b, config.CurrentConfig.Welcome.Message, &gotgbot.SendMessageOpts{})
		return err
	}

	return nil
}
