package _interface

import (
	"github.com/bwmarrin/discordgo"
	"gitlab.com/one-more/acropole/app/service"
)

type Interface struct {
	Discord *discordgo.Session
	Service *service.Service
}

type Handler interface {
	addHandler(discord *discordgo.Session)
}

func (i *Interface) InitInterfaces() error {
	i.addPollHandlers()

	return nil
}
