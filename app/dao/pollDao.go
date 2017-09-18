package dao

import (
	"gitlab.com/one-more/acropole/app/model"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

const database = "acropole"
const pollCollection = "poll"

func (dao *Dao) AddPoll(poll *model.Poll) error {
	fmt.Println("Insert poll")
	fmt.Println(poll)
	return dao.db.session.DB(database).C(pollCollection).Insert(poll)
}

func (dao *Dao) Vote(pollId string, vote *model.Vote) error {
	query := bson.M{
		"_id":           bson.ObjectIdHex(pollId),
		"votes.user_id": bson.M{"$ne": vote.UserId}, // Avoid two vote per user
	}
	update := bson.M{"$addToSet": bson.M{"votes": vote}}

	return dao.db.session.DB(database).C(pollCollection).Update(query, update)
}
