package commands

import "github.com/bwmarrin/discordgo"

var SexyTime = Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "sexy-time",
		Description: "YOU KNOW WHAT TIME IT IS!!! 👍 👍",
	},
	Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "https://tenor.com/view/sexytime-sexy-borat-very-nice-gif-18791202",
			},
		})
	},
}
