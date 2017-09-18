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
	configInstance := config.InitConfig()

	/*
	 * DAO layer initialisation
	 */

	daoInstance, err := dao.NewDao(configInstance.Mongodb)

	if err != nil {
		panic(err)
	}

	defer daoInstance.Close()

	/*
	 * Service layer initialisation
	 */

	serviceInstance := &service.Service{Dao: daoInstance}

	/*
	 * Interface layer initialisation
	 */

	var token string = "Bot " + configInstance.Discord.Token // TODO: handle user/password connection to Discord
	discord, err := discordgo.New(token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	discordInterface := _interface.Interface{Discord: discord, Service: serviceInstance}

	if err = discordInterface.InitInterfaces(); err != nil {
		fmt.Println("error while interfacing with discord,", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	if err = discord.Open(); err != nil {
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