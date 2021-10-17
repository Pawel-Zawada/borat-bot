package commands

import (
	"github.com/bwmarrin/discordgo"
	"nextrock/borat_bot/discord/voice"
)

var SexyTime = Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "sexy-time",
		Description: "YOU KNOW WHAT TIME IT IS!!! 👍 👍",
	},
	Handler: func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		go session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "https://tenor.com/view/sexytime-sexy-borat-very-nice-gif-18791202",
			},
		})

		go voice.PlaySound(session, interaction.GuildID, interaction.Member.User.ID, "very-nice.mp3")
	},
}
