package estralserver

import (
	"fmt"
	"log"
	"os"

	"github.com/EstralMC/GoMine/estral/console"
	"github.com/EstralMC/GoMine/estral/handlers"
	"github.com/EstralMC/GoMine/server"
	"github.com/EstralMC/GoMine/server/player"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

var Config, err = readConfig(console.Log)
var Srv = Config.New()

func Start() {
	console.Log.Formatter = &logrus.TextFormatter{ForceColors: true}
	console.Log.Level = logrus.DebugLevel

	if err != nil {
		log.Fatalln(err)
	}

	Srv.CloseOnProgramEnd()

	go console.Console()

	Srv.Listen()
	for Srv.Accept(func(p *player.Player) {
		p.Handle(handlers.New(p))

		console.Src.SendMessagef("§e%v has connected to the server!", p.Name())
	}) {
	}
}

func readConfig(log server.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return c.Config(log)
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c.Config(log)
}
