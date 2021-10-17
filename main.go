package main

import (
	"log"
	"nextrock/borat_bot/discord"
	"os"
	"os/signal"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go discord.Run()

	// Keep running until received shutdown signal
	sig := <-stop
	discord.Stop()
	log.Println("Stopping Go process", sig)
}
