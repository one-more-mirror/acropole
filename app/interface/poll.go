package _interface

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"fmt"
	"log"
	"regexp"
	"gitlab.com/one-more/acropole/app"
)

type Handler interface {
	Handle() func(*discordgo.Session, *discordgo.MessageCreate)
}

type BanPollHandler struct {
	Interface
}

func (i *BanPollHandler) Handle() func(*discordgo.Session, *discordgo.MessageCreate) {

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {

		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		if !strings.HasPrefix(m.Content, "!democracy ban") {
			// TODO: print usage
			return
		}

		if len(m.Mentions) != 1 {
			return
		}

		// User to ban
		user := m.Mentions[0]

		channel, err := s.Channel(m.ChannelID)

		if err != nil {
			// TODO: log error
			return
		}

		newPoll := acropole.NewBanPoll(m.Author.ID, &acropole.BanPollPayload{
			UserId:  user.ID,
			GuildId: channel.GuildID,
		})

		if err = i.PollService.CreatePoll(newPoll); err != nil {
			// TODO: log error
			return
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Ban poll #%s created without error :ok_hand:", newPoll.Id))
	}
}

type KickPollHandler struct {
	Interface
}

func (i *KickPollHandler) Handle() func(*discordgo.Session, *discordgo.MessageCreate) {

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {

		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		if !strings.HasPrefix(m.Content, "!democracy kick") {
			return
		}

		if len(m.Mentions) != 1 {
			// TODO: print usage
			return
		}

		// User to ban
		user := m.Mentions[0]

		channel, err := s.Channel(m.ChannelID)
		if err != nil {
			// TODO: log error
			return
		}

		// New poll
		newPoll := acropole.NewKickPoll(m.Author.ID, &acropole.KickPollPayload{
			UserId:  user.ID,
			GuildId: channel.GuildID,
		})

		if err = i.PollService.CreatePoll(newPoll); err != nil {
			// TODO: log error
			return
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Kick poll #%s created without error :ok_hand:", newPoll.Id))
	}
}

type VoteHandler struct {
	Interface
}

func (i *VoteHandler) Handle() func(*discordgo.Session, *discordgo.MessageCreate) {

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {

		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		if !strings.HasPrefix(m.Content, "!democracy vote") {
			return
		}

		r, _ := regexp.Compile("^!democracy vote ([0-9a-f]{24}) (yes|no)$")

		args := r.FindStringSubmatch(m.Content)

		if len(args) != 3 {
			// TODO: print usage
			return
		}

		pollId := args[1]
		voteArg := args[2]

		vote := &acropole.Vote{
			UserId: m.Author.ID,
			Time:   m.Timestamp,
			Yes:    voteArg == "yes",
		}

		if err := i.PollService.CreateVote(pollId, vote); err != nil {
			// TODO: log error
			log.Println(fmt.Sprintf("%#v", err))
			return
		}

		s.ChannelMessageSend(m.ChannelID, "Voted without error :ok_hand:")

	}
}
