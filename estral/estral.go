package estral

import (
	"github.com/EstralMC/GoMine/estral/commands"
	"github.com/EstralMC/GoMine/estral/entities/register"
	"github.com/EstralMC/GoMine/estral/estralserver"
	"github.com/EstralMC/GoMine/server/cmd"
)

func Start() {
	register.RegisterEntities()
	registerCommands()

	estralserver.Start()
}

func registerCommands() {
	cmd.Register(cmd.New("gamemode", "Changes game mode", []string{}, commands.GamemodeCmd{}))
	cmd.Register(cmd.New("op", "Grants Operator Access", []string{}, commands.OpCmd{}))
	cmd.Register(cmd.New("deop", "Takes away Operator Access", []string{}, commands.DeopCmd{}))
	cmd.Register(cmd.New("say", "Broadcasts message", []string{}, commands.SayCmd{}))
	cmd.Register(cmd.New("summon", "Summons an entity", []string{}, commands.Summon{}))
	cmd.Register(cmd.New("weather", "Sets the world weather", []string{}, commands.Weather{}))
}
