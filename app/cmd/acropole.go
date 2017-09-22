package main

import (
	"gopkg.in/mgo.v2"
	"gitlab.com/one-more/acropole/app/db"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/op/go-logging"
	"gitlab.com/one-more/acropole/app"
	"gitlab.com/one-more/acropole/app/config"
	"gitlab.com/one-more/acropole/app/interface"
)

var log = logging.MustGetLogger("acropole")

func main() {
	/*
	 * Configuration initialisation
	 */
	configInstance := config.InitConfig()

	/*
	 * DAO layer initialisation
	 */

	session, err := mgo.Dial(configInstance.Mongodb.Host)

	if err != nil {
		panic(err)
	}

	defer session.Close()

	/*
	 * Interface layer initialisation
	 */

	var token string = "Bot " + configInstance.Discord.Token // TODO: handle user/password connection to Discord
	discord, err := discordgo.New(token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	if err = discord.Open(); err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	defer discord.Close()

	var ps acropole.PollService
	ps = acropole.PollService(&db.PollService{Session: session})

	_, err = _interface.NewInterface(discord, ps)
	if err != nil {
		fmt.Println("error while interfacing with discord,", err)
		return
	}

	log.Info("Bot is now running!")

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
