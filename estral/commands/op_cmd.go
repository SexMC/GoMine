package commands

import (
	"github.com/EstralMC/GoMine/estral/memory"
	"github.com/EstralMC/GoMine/server/cmd"
	"github.com/EstralMC/GoMine/server/player"
)

type OpCmd struct {
	Target []cmd.Target `cmd:"player"`
}

func (c OpCmd) Run(src cmd.Source, o *cmd.Output) {
	s, isPlayer := src.(*player.Player)

	if len(c.Target) > 1 {
		o.Print("§cYou cannot use this command on more than 1 players at a time!")
		return
	}

	t := c.Target[0].(*player.Player)

	if memory.OperatorExists(t.UUID()) {
		o.Printf("§7Nothing changed! §f%v §7is already a Server Operator", t.Name())
		return
	}

	memory.AddOperator(t.UUID())
	t.SendToast("§2Operator Access Granted", "You are now a §eServer Operator!")

	if isPlayer && s.UUID() == t.UUID() {
		return
	}

	o.Printf("§7You have made §f" + t.Name() + " §7a Server Operator")
}
