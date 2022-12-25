package commands

import (
	"github.com/EstralMC/GoMine/server/cmd"
	"github.com/EstralMC/GoMine/server/player"
	"github.com/EstralMC/GoMine/server/world"
)

type GamemodeCmd struct {
	Mode   Gamemode                   `cmd:"gamemode"`
	Target cmd.Optional[[]cmd.Target] `cmd:"player"`
}

func (c GamemodeCmd) Run(src cmd.Source, o *cmd.Output) {
	players, _ := c.Target.Load()

	s, isPlayer := src.(*player.Player)

	if len(players) == 0 {
		if isPlayer {
			_ = UpdateGamemode(s, c.Mode)
		} else {
			o.Errorf("Usage: /gamemode <mode> <player>")
		}
	}

	for _, target := range players {
		t := target.(*player.Player)
		mode := UpdateGamemode(t, c.Mode)

		if len(players) < 5 {
			if isPlayer && t.UUID() == s.UUID() {
				return
			}

			o.Printf("§7You have updated §f%v's §7Gamemode to §f%v", t.Name(), mode)
			return
		}

		o.Printf("§7You have updated §f%v Player's §7Gamemode to §f%v", len(players), mode)
	}
}

func UpdateGamemode(t *player.Player, Mode Gamemode) string {
	var mode = ""

	switch Mode {
	case "s", "survival", "0":
		mode = "Survival"
		t.SetGameMode(world.GameModeSurvival)
		break
	case "c", "creative", "1":
		mode = "Creative"
		t.SetGameMode(world.GameModeCreative)
		break
	case "a", "adventure", "2":
		mode = "Adventure"
		t.SetGameMode(world.GameModeAdventure)
		break
	case "sp", "spectator", "3":
		mode = "Spectator"
		t.SetGameMode(world.GameModeSpectator)
		break
	}

	t.Messagef("§7Your Game Mode has been updated to §f%v", mode)
	return mode
}

type Gamemode string

func (Gamemode) Type() string {
	return "gamemode"
}

func (Gamemode) Options(_ cmd.Source) []string {
	return []string{"s", "survival", "0", "c", "creative", "1", "a", "adventure", "2", "sp", "spectator", "3"}
}
