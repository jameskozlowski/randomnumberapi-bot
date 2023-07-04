package main

import (
	"log"
	"os"
)

func main() {
	// Load environment variables
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable: BOT_TOKEN")
	}

	// Start the bot
	bot := RandomnumberapiDiscordBot{
		BotToken: token,
	}

	bot.Run()
}
