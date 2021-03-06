package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"nextrock/borat_bot/commands"
	"os"
)

var Discord *discordgo.Session

// create a bot configuration that is ready to be connected
func create() {
	var err error
	Discord, err = discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	log.Println("Great success!")
}

// connect to Discord websocket
func connect() {
	var err error
	err = Discord.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
}

// handleReady sets a handler for when the bot has successfully connected
func handleReady() chan *discordgo.Ready {
	channel := make(chan *discordgo.Ready)

	Discord.AddHandler(func(session *discordgo.Session, ready *discordgo.Ready) {
		log.Println("Bot is up!")
		channel <- ready
	})

	return channel
}

// initCommands send command creation requests for each defined command in commands.Commands
// and assign corresponding handlers for each.
func initCommands() {
	for _, v := range commands.Commands {
		go func(v commands.Command) {
			_, err := Discord.ApplicationCommandCreate(Discord.State.User.ID, "", v.ApplicationCommand)
			if err != nil {
				log.Panicf("Cannot create '%v' command: %v", v.ApplicationCommand.Name, err)
			}
			log.Printf("Loaded command: %v", v.ApplicationCommand.Name)

			Discord.AddHandler(func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
				v.Handler(session, interaction)
			})
		}(v)
	}
}

// Run starts up the Discord bot connection and loads in the application commands
func Run() {
	create()

	ready := handleReady()
	connect()

	<-ready
	initCommands()
}

// Stop gracefully shuts down the Discord connection
func Stop() {
	log.Println("Gracefully shutdowning Discord connection")
	err := Discord.Close()
	if err != nil {
		log.Panicf("Failed to gracefully disconnect from Discord, error: '%v'", err)
	}
}
