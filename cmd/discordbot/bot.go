package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jameskozlowski/randomnumberapi-bot/internal/randomnumberapiclient"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

type RandomnumberapiDiscordBot struct {
	BotToken string
}

func (bot *RandomnumberapiDiscordBot) Run() {
	// Create new Discord Session
	discord, err := discordgo.New("Bot " + bot.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	// Add event handler
	discord.AddHandler(bot.newMessage)

	// Open session
	discord.Open()
	defer discord.Close()

	// Run until code is terminated
	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func (bot *RandomnumberapiDiscordBot) newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messaage
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Respond to messages
	switch {
	case strings.Compare(message.Content, "!random") == 0:
		random, err := randomnumberapiclient.GetRandomNumber()
		if err != nil {
			fmt.Println(err)
			msg := message.Author.Username + ":  Internal Error getting number"
			discord.ChannelMessageSend(message.ChannelID, msg)
		} else {
			msg := message.Author.Username + ":  your random number is:  " + strconv.Itoa(random)
			discord.ChannelMessageSend(message.ChannelID, msg)
		}
	case strings.Compare(message.Content, "!random reddit") == 0:
		random, err := randomnumberapiclient.GetRandomRedditNumber()
		if err != nil {
			fmt.Println(err)
			msg := message.Author.Username + ":  Internal Error getting number"
			discord.ChannelMessageSend(message.ChannelID, msg)
		} else {
			msg := message.Author.Username + ":  your random reddit number is:  " + strconv.Itoa(random)
			discord.ChannelMessageSend(message.ChannelID, msg)
		}
	case strings.Contains(message.Content, "!random"):
		discord.ChannelMessageSend(message.ChannelID, "you can use\n\n"+
			"'!random' to get a random number between 1 and 100\n\n"+
			"'!random min max' to get a random number between min and max\n\n"+
			"'!random reddit to get a random number between 1 and 100 using a reddit comment as a seed\n\n"+
			"'!random reddit min max' to get a random number between min and max using a reddit comment as a seed")
	}
}
