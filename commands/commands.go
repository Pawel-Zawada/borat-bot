package commands

import (
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	ApplicationCommand *discordgo.ApplicationCommand
	Handler            func(*discordgo.Session, *discordgo.InteractionCreate)
}

var Commands = map[string]Command{
	SexyTime.ApplicationCommand.Name: SexyTime,
}
