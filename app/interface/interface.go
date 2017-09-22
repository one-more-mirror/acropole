package _interface

import (
	"github.com/bwmarrin/discordgo"
	"gitlab.com/one-more/acropole/app"
)

var handlers = []Handler{
	&BanPollHandler{},
	&KickPollHandler{},
	&VoteHandler{},
}

type Interface struct {
	Discord     *discordgo.Session
	PollService acropole.PollService
}

func NewInterface(s *discordgo.Session, ps acropole.PollService) (Interface, error) {
	i := Interface{
		Discord:     s,
		PollService: ps,
	}

	for _, h := range handlers {
		i.Discord.AddHandler(h.Handle())
	}

	return i, nil
}
