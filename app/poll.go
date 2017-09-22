package acropole

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/bwmarrin/discordgo"
)

type PollService interface {
	Poll(id string) (*Poll, error)
	Polls() (*[]Poll, error)
	CreatePoll(poll *Poll) error
	DeletePoll(id string) error

	Votes(pollId string) error
	CreateVote(pollId string, vote *Vote) error
	DeleteVote(pollId string, userId string) error
}

type Poll struct {
	Id        string      `json:"id" bson:"_id,omitempty"`
	Votes     []*Vote     `json:"votes" bson:"votes"`
	CreatorId string      `json:"creator_id" bson:"creator_id"`
	PollType  string      `json:"type" bson:"type"`
	Payload   interface{} `json:"payload" bson:"payload"`
}

type Vote struct {
	Time   discordgo.Timestamp `json:"time" bson:"time"`
	UserId string              `json:"user_id" bson:"user_id"`
	Yes    bool                `json:"yes" bson:"yes"`
}

type BanPollPayload struct {
	UserId  string `json:"user_id" bson:"user_id"`
	GuildId string `json:"guild_id" bson:"guild_id"`
}

type KickPollPayload struct {
	UserId  string `json:"user_id" bson:"user_id"`
	GuildId string `json:"guild_id" bson:"guild_id"`
}

func newPoll(creatorId string, pollType string, payload interface{}) *Poll {
	return &Poll{
		Id:        bson.NewObjectId().Hex(),
		CreatorId: creatorId,
		PollType:  pollType,
		Payload:   payload,
	}
}

func NewBanPoll(creatorId string, payload *BanPollPayload) *Poll {
	return newPoll(creatorId, "ban", payload)
}

func NewKickPoll(creatorId string, payload *KickPollPayload) *Poll {
	return newPoll(creatorId, "kick", payload)
}
