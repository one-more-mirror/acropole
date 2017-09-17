package dao

import (
	"gopkg.in/mgo.v2"
	model "../model"
)

type Dao struct {
	session *mgo.Session
}

func (dao *Dao) GetPools() ([]*model.Poll, error) {
	pools := []*model.Poll{
		{
			Votes: []*model.Vote{
				{
					Time: nil,
					User: nil,
					Yes:  true,
				},
				{
					Time: nil,
					User: nil,
					Yes:  false,
				},
			},
			Creator: nil,
			Id:      1,
		},
	}

	return pools, nil
}
