package main

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"gitlab.com/one-more/acropole/app/dao"
	"gitlab.com/one-more/acropole/app/config"
	"gitlab.com/one-more/acropole/app/interface"
	"gitlab.com/one-more/acropole/app/service"
)

func main() {

	// Initialize configuration
	config := config.InitConfig()

	// Database connection initialisation
	dao, err := dao.NewDao(config.Mongodb)

	if err != nil {
		panic(err)
	}

	defer dao.Close()

	// Service initialisation
	serviceInstance := &service.Service{Dao: dao}

	// Connect to Discord
	var token string = "Bot " + config.Discord.Token // TODO: handle user/password connection to Discord
	discord, err := discordgo.New(token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	discordInterface := _interface.Interface{Discord: discord, Service: serviceInstance}
	err = discordInterface.InitInterfaces()

	if err != nil {
		fmt.Println("error while interfacing with discord,", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	defer discord.Close()

	fmt.Println("Bot is now running!")

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

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
