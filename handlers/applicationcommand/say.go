package applicationcommand

import (
	"log/slog"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

// AutocompleteInteractionCreate
var say = bot.NewListenerFunc(func(event *events.ApplicationCommandInteractionCreate) {
	if event.SlashCommandInteractionData().CommandName() != "say" {
		return
	}

	data := event.SlashCommandInteractionData()
	err := event.CreateMessage(discord.NewMessageCreateBuilder().
		SetContent(data.String("message")).
		SetEphemeral(data.Bool("ephemeral")).
		Build(),
	)
	if err != nil {
		slog.Error("failed to send message", slog.Any("err", err))
	}
})
