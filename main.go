package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		log.Fatal("Failed to load .env file: ", dotenvErr.Error())
	}

	var (
		appID        = os.Getenv("APP_ID")
		guildID      = os.Getenv("GUILD_ID")
		discordToken = os.Getenv("DISCORD_TOKEN")
	)

	discordSession, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err.Error())
	}

	// Connects to Database
	// TODO: pass DB down to commands to run queries
	// db := database.Init()
	// defer db.Close()

	loadCommands(discordSession, appID, guildID)

	// Opening the discord session
	err = discordSession.Open()
	if err != nil {
		log.Fatal("Error opening Discord session: ", err.Error())
	}

	fmt.Println("Guild bot is running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discordSession.Close()
}
