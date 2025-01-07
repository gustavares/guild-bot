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

	// List of bot commands
	_, acErr := discordSession.ApplicationCommandBulkOverwrite(appID, guildID, []*discordgo.ApplicationCommand{
		{
			Name:        "hello-world",
			Description: "Showcase of a basic slash command",
		},
	})
	if acErr != nil {
		log.Fatal("Failed to set bot commands: ", err.Error())
	}

	discordSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.ApplicationCommandData()
		switch data.Name {
		case "hello-world":
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Hello world!",
					},
				},
			)
			if err != nil {
				fmt.Println("hello-world command failed: ", err.Error())
			}
		}
	})

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
