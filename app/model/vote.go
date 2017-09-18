package model

import "github.com/bwmarrin/discordgo"

type Vote struct {
	Time   discordgo.Timestamp `json:"time" bson:"time"`
	UserId string              `json:"user_id" bson:"user_id"`
	Yes    bool                `json:"yes" bson:"yes"`
}
