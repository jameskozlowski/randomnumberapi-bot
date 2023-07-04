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

	// Run until terminated
	fmt.Println("---Running randomnumberapi bot---")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func (bot *RandomnumberapiDiscordBot) newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messages
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Respond to messages
	switch {
	case strings.HasPrefix(message.Content, "!random reddit"):
		min, max := getMinMax(message.Content)
		random, err := randomnumberapiclient.GetRandomRedditNumber(min, max)
		if err != nil {
			fmt.Println(err)
			msg := message.Author.Username + ":  Internal Error getting number"
			discord.ChannelMessageSend(message.ChannelID, msg)
		} else {
			msg := message.Author.Username + ":  your random reddit number is:  " + strconv.Itoa(random)
			discord.ChannelMessageSend(message.ChannelID, msg)
		}
	case strings.HasPrefix(message.Content, "!random"):
		min, max := getMinMax(message.Content)
		random, err := randomnumberapiclient.GetRandomNumber(min, max)
		if err != nil {
			fmt.Println(err)
			msg := message.Author.Username + ":  Internal Error getting number"
			discord.ChannelMessageSend(message.ChannelID, msg)
		} else {
			msg := message.Author.Username + ":  your random number is:  " + strconv.Itoa(random)
			discord.ChannelMessageSend(message.ChannelID, msg)
		}
	case strings.HasPrefix(message.Content, "!random help"):
		discord.ChannelMessageSend(message.ChannelID, "you can use\n\n"+
			"'!random' to get a random number between 1 and 100\n\n"+
			"'!random min max' to get a random number between min and max\n\n"+
			"'!random reddit to get a random number between 1 and 100 using a reddit comment as a seed\n\n"+
			"'!random reddit min max' to get a random number between min and max using a reddit comment as a seed")
	}
}

func getMinMax(msg string) (int, int) {
	split := strings.Split(msg, " ")
	if len(split) != 3 && len(split) != 4 {
		return 0, 100
	}
	//set the start index based on if !random or !random reddit
	minStartPosition := len(split) - 3
	min, err := strconv.Atoi(split[minStartPosition+1])
	if err != nil || min < 0 {
		return 0, 100
	}
	max, err := strconv.Atoi(split[minStartPosition+2])
	if err != nil || max <= min {
		return 0, 100
	}
	return min, max
}
