package model

type Poll struct {
	Id        string      `json:"id" bson:"_id,omitempty"`
	Votes     []*Vote     `json:"votes" bson:"votes"`
	CreatorId string      `json:"creator_id" bson:"creator_id"`
	PollType  string      `json:"type" bson:"type"`
	Payload   interface{} `json:"payload" bson:"payload"`
}

type BanPollPayload struct {
	UserId  string `json:"user_id" bson:"user_id"`
	GuildId string `json:"guild_id" bson:"guild_id"`
}

func NewBanPoll(creatorId string, payload *BanPollPayload) *Poll {
	return &Poll{
		CreatorId: creatorId,
		PollType:  "ban",
		Payload:   payload,
	}
}

type kickPoll struct {
	Poll
	UserId  string
	GuildId string
}
