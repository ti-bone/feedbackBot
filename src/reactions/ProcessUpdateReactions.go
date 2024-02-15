/*
 * ProcessUpdateReactions.go
 * Copyright (c) ti-bone 2023-2024
 */

package reactions

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"log"
	"os"
)

func ProcessUpdateReactions(
	reaction *gotgbot.MessageReactionUpdated, b *gotgbot.Bot,
	chatId int64, messageId int64,
) error {
	// If some reactions removed - remove them from the message by calling same method once again
	if len(reaction.OldReaction) != 0 {
		_, err := b.SetMessageReaction(
			chatId, messageId,
			&gotgbot.SetMessageReactionOpts{
				Reaction: reaction.OldReaction,
			},
		)

		// Errors may occur here, but they aren't critical, just log them
		if err != nil {
			log.SetOutput(os.Stderr)
			log.Println(err.Error())
		}
	}

	// If the message is found - send the reaction to the user
	_, err := b.SetMessageReaction(
		chatId, messageId,
		&gotgbot.SetMessageReactionOpts{
			Reaction: reaction.NewReaction,
		},
	)

	return err
}
