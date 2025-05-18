package applicationcommand

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var feedback = bot.NewListenerFunc(func(event *events.ApplicationCommandInteractionCreate) {
	if event.SlashCommandInteractionData().CommandName() != "feedback" {
		return
	}

	data := event.SlashCommandInteractionData()
	err := event.CreateMessage(discord.NewMessageCreateBuilder().
		SetContent(fmt.Sprintf("You have left the following %s feedback for %s\n%s",
			strings.ToLower(data.String("type")),
			data.String("name"),
			data.String("feedback"),
		)).
		SetEphemeral(true).
		Build(),
	)
	if err != nil {
		slog.Error("failed to send message", slog.Any("err", err))
	}
})
