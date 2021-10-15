package commands

import (
	"regexp"
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestCreateCommand(t *testing.T) {
	command := Command{
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "test-command",
			Description: "This is a test description",
		},
		Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		},
	}
	match := "test-command"
	want := regexp.MustCompile(`\b` + command.ApplicationCommand.Name + `\b`)
	if !want.MatchString(match) {
		t.Fatalf(`Command.ApplicationCommand.Name = %q, want match for %#q`, command.ApplicationCommand.Name, match)
	}
}
