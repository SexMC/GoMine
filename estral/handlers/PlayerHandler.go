package handlers

import (
	"github.com/EstralMC/GoMine/estral/console"
	"github.com/EstralMC/GoMine/server/event"
	"github.com/EstralMC/GoMine/server/player"
)

type PlayerHandler struct {
	player.NopHandler
	p *player.Player
}

func New(p *player.Player) *PlayerHandler {
	return &PlayerHandler{p: p}
}

func (m *PlayerHandler) HandleChat(ctx *event.Context, message *string) {
	console.Src.SendMessagef("%s: %s", m.p.Name(), *message)
}

func (m *PlayerHandler) HandleQuit() {
	console.Src.SendMessagef("§e%v has disconnected from the server!", m.p.Name())
}
