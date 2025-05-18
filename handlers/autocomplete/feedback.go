package autocomplete

import (
	"log/slog"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var feedback = bot.NewListenerFunc(func(event *events.AutocompleteInteractionCreate) {
	data := event.Data
	if data.CommandName != "feedback" {
		return
	}

	focused := data.Focused()
	if focused.Name != "name" {
		return
	}

	value := data.String("name")
	if len(value) < 1 {
		return
	}

	if err := event.AutocompleteResult([]discord.AutocompleteChoice{
		discord.AutocompleteChoiceString{
			Name:  "Choice 1",
			Value: "Choice 1",
		},
		discord.AutocompleteChoiceString{
			Name:  "Choice 2",
			Value: "Choice 2",
		},
		discord.AutocompleteChoiceString{
			Name:  "Choice 3",
			Value: "Choice 3",
		},
	}); err != nil {
		slog.Error("failed to respond to autocomplete interaction", slog.Any("err", err))
	}
})
