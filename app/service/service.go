package service

import (
	"gitlab.com/one-more/acropole/app/dao"
	"gitlab.com/one-more/acropole/app/model"
)

type Service struct {
	Dao *dao.Dao
}

func (s *Service) AddPoll(poll *model.Poll) error {
	return s.Dao.AddPoll(poll)
}

func (s *Service) Vote(pollId string, vote *model.Vote) error {
	return s.Dao.Vote(pollId, vote)
}

