package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/EstralMC/GoMine/server/cmd"
	"github.com/EstralMC/GoMine/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()
var Src = Source{log: Log}

func Console() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if t := strings.TrimSpace(scanner.Text()); len(t) > 0 {
			name := strings.Split(t, " ")[0]
			if c, ok := cmd.ByAlias(name); ok {
				c.Execute(strings.TrimPrefix(strings.TrimPrefix(t, name), " "), Src)
			} else {
				output := &cmd.Output{}
				output.Errorf("Could not find command '%s'", name)
				Src.SendCommandOutput(output)
			}
		}
	}
}

// Source is the command estral used to execute commands from the console.
type Source struct {
	log *logrus.Logger
}

// Name returns the name of console.
func (Source) Name() string { return "CONSOLE" }

// Position ...
func (Source) Position() mgl64.Vec3 { return mgl64.Vec3{} }

// SendCommandOutput prints out command outputs.
func (s Source) SendCommandOutput(o *cmd.Output) {
	for _, e := range o.Errors() {
		s.log.Error(text.ANSI(e))
	}
	for _, m := range o.Messages() {
		s.log.Info(text.ANSI(m))
	}
}

// SendMessage prints out message in console
func (s Source) SendMessage(message string) {
	message = format(fmt.Sprintln(message))
	s.log.Info(text.ANSI(message + "§r"))
}

// SendMessagef sends a formatted message using a specified format to console
func (s Source) SendMessagef(message string, args ...any) {
	s.SendMessage(fmt.Sprintf(message, args...))
}

// World ...
func (Source) World() *world.World { return nil }

// format is a utility function to format a list of values to have spaces between them, but no newline at the
// end, which is typically used for sending messages, popups and tips.
func format(text string) string {
	return strings.TrimSuffix(strings.TrimSuffix(text, "\n"), "\n")
}
