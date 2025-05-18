package applicationcommand

import "github.com/disgoorg/disgo/bot"

var Handlers = []bot.EventListener{
	feedback,
	say,
}
