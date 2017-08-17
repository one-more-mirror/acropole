package main

import (
	"github.com/spf13/viper"
	"github.com/bwmarrin/discordgo"
	"github.com/kr/pretty"
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type mongoConfig struct {
	Username string
	Password string
	Host     string
	Port     int
}

type discordConfig struct {
	Token    string
	Username string
	Password string
}

type config struct {
	Mongodb mongoConfig
	Discord discordConfig
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.AllSettings()
	var config config

	err = viper.Unmarshal(&config)

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Invalid config file: %s \n", err))
	}

	session, err := mgo.Dial(config.Mongodb.Host)

	if err != nil {
		panic(err)
	}

	defer session.Close()

	var token string = "Bot " + config.Discord.Token

	discord, err := discordgo.New(token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	discord.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running!  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong local!!!!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping! :)")
	}
}