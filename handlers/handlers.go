package handlers

import (
	"go-pro-bot/handlers/applicationcommand"
	"go-pro-bot/handlers/autocomplete"
)

var Handlers = append(applicationcommand.Handlers, autocomplete.Handlers...)
