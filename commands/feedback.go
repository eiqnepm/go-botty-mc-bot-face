package commands

import (
	"github.com/disgoorg/disgo/discord"
)

var feedback = discord.SlashCommandCreate{
	Name:        "feedback",
	Description: "leave a feedback",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name:         "name",
			Description:  "Name of recipient",
			Required:     true,
			Autocomplete: true,
		},
		discord.ApplicationCommandOptionString{
			Name:        "type",
			Description: "Type of feedback",
			Required:    true,
			Choices: []discord.ApplicationCommandOptionChoiceString{
				{Name: "Positive", Value: "Positive"},
				{Name: "Neutral", Value: "Neutral"},
				{Name: "Negative", Value: "Negative"},
			},
		},
		discord.ApplicationCommandOptionString{
			Name:        "feedback",
			Description: "Feedback",
			Required:    true,
		},
	},
}
