package commands

import (
	"feedbackBot/src/db"
	"feedbackBot/src/helpers"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Unban(b *gotgbot.Bot, ctx *ext.Context) error {
	args := ctx.Args()

	if len(args) <= 1 {
		_, err := ctx.EffectiveMessage.Reply(b, "-400: user not specified.", &gotgbot.SendMessageOpts{})
		return err
	}

	user, err := helpers.ParseInputUser(args[1])

	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
		return err
	}

	if !user.IsBanned {
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("#u%d is not banned.", user.UserID), &gotgbot.SendMessageOpts{})
		return err
	}

	db.Connection.Model(&user).Update("is_banned", false)

	_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Ban lifted from #u%d.", user.UserID), &gotgbot.SendMessageOpts{})
	return err
}
