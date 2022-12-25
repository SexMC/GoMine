package commands

import (
	"github.com/EstralMC/GoMine/estral/utils"
	"github.com/EstralMC/GoMine/server/cmd"
	"github.com/EstralMC/GoMine/server/player"
)

type SayCmd struct {
	Message cmd.Varargs `cmd:"message"`
}

func (c SayCmd) Run(src cmd.Source, o *cmd.Output) {
	p, isPlayer := src.(*player.Player)

	sender := ""
	if isPlayer {
		sender = p.Name()
	} else {
		sender = "Console"
	}

	utils.BroadcastMessagef("§d%s says: %s", sender, c.Message)
}
