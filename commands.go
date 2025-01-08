package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func helloWorldCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

func loadCommands(s *discordgo.Session, appID, guildID string) {
	_, err := s.ApplicationCommandBulkOverwrite(appID, guildID, []*discordgo.ApplicationCommand{
		{
			Name:        "hello-world",
			Description: "Showcase of a basic slash command",
		},
	})
	if err != nil {
		log.Fatal("Failed to set bot commands: ", err.Error())
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.ApplicationCommandData()
		switch data.Name {
		case "hello-world":
			helloWorldCommand(s, i)
		}
	})
}
