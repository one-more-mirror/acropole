package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gitlab.com/one-more/acropole/app"
)

const database = "acropole"
const pollCollection = "poll"

// MongoDB implementation of acropole.PollService
type PollService struct {
	Session *mgo.Session
}

func (ps *PollService) Poll(id string) (*acropole.Poll, error) {
	panic("implement me")
}

func (ps *PollService) Polls() (*[]acropole.Poll, error) {
	panic("implement me")
}

func (ps *PollService) CreatePoll(poll *acropole.Poll) error {
	return ps.Session.DB(database).C(pollCollection).Insert(poll)
}

func (ps *PollService) DeletePoll(id string) error {
	panic("implement me")
}

func (ps *PollService) Votes(pollId string) error {
	panic("implement me")
}

func (ps *PollService) CreateVote(pollId string, vote *acropole.Vote) error {
	query := bson.M{
		"_id":           pollId,
		"votes.user_id": bson.M{"$ne": vote.UserId}, // Avoid two vote per user
	}

	update := bson.M{"$addToSet": bson.M{"votes": vote}}

	return ps.Session.DB(database).C(pollCollection).Update(query, update)
}

func (ps *PollService) DeleteVote(pollId string, userId string) error {
	panic("implement me")
}
