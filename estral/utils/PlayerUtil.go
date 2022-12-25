package utils

import (
	"fmt"

	"github.com/EstralMC/GoMine/estral/console"
	"github.com/EstralMC/GoMine/estral/estralserver"
)

// BroadcastMessage broadcasts a message to all players and the console
func BroadcastMessage(message string) {
	message = fmt.Sprintln(message)

	players := estralserver.Srv.Players()

	for _, player := range players {
		player.Message(message)
	}

	console.Src.SendMessage(message)
}

// BroadcastMessagef broadcasts a formatted message to all the players and the console
func BroadcastMessagef(message string, args ...any) {
	BroadcastMessage(fmt.Sprintf(message, args...))
}
