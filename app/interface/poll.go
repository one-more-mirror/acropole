package _interface

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"gitlab.com/one-more/acropole/app/model"
	"encoding/json"
)

func (i *Interface) addPollHandlers() {
	i.Discord.AddHandler(i.banPollHandler())
	i.Discord.AddHandler(i.voteHandler())
}

func (i *Interface) banPollHandler() func(*discordgo.Session, *discordgo.MessageCreate) {

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {

		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		if strings.HasPrefix(m.Content, "!democracy ban") {
			channel, err := s.Channel(m.ChannelID)

			if err != nil {
				msg, _ := json.Marshal(err)
				s.ChannelMessageSend(m.ChannelID, "Error: "+string(msg))
			}

			newPoll := model.NewBanPoll(m.Author.ID, &model.BanPollPayload{
				UserId:  m.Author.ID,
				GuildId: channel.GuildID,
			})

			err = i.Service.AddPoll(newPoll)

			if err != nil {
				msg, _ := json.Marshal(err)
				s.ChannelMessageSend(m.ChannelID, "Error: "+string(msg))
			}
			s.ChannelMessageSend(m.ChannelID, "Ban poll created without error :ok_hand:")
		}
	}
}

func (i *Interface) voteHandler() func(*discordgo.Session, *discordgo.MessageCreate) {

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {

		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		if strings.HasPrefix(m.Content, "!democracy vote") {

			vote := &model.Vote{
				UserId: m.Author.ID,
				Time:   m.Timestamp,
				Yes:    true,
			}

			err := i.Service.Vote("59c023ca1632b5715bfa2f5b", vote)

			if err != nil {
				msg, _ := json.Marshal(err)
				s.ChannelMessageSend(m.ChannelID, "Error: "+string(msg))
			}

			s.ChannelMessageSend(m.ChannelID, "Voted without error :ok_hand:")
		}
	}
}
