package commands

import (
	"github.com/EstralMC/GoMine/estral/entities/types"
	"github.com/EstralMC/GoMine/server/cmd"
	"github.com/EstralMC/GoMine/server/player"
)

type Summon struct {
	Mob  string `cmd:"mob"`
	Name string `cmd:"nametag"`
}

func (c Summon) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)

	z := types.GetMobFromType(c.Name, p.Position(), true, c.Mob)

	if z == nil {
		o.Printf("§cError: Mob Type %s is unimplemented", c.Mob)
		return
	}

	p.World().AddEntity(z)
	o.Printf("§bMob Spawned with the name: %s", c.Name)
}
