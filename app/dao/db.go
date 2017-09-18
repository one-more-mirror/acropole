package dao

import (
	"gopkg.in/mgo.v2"
	"gitlab.com/one-more/acropole/app/config"
)

type Db struct {
	config  config.MongoConfig
	session *mgo.Session
}

func newDb(config config.MongoConfig) (*Db, error) {
	session, err := mgo.Dial(config.Host)

	if err != nil {
		return nil, err
	}

	db := &Db{config: config, session: session}

	return db, nil
}

func (db *Db) Close() {
	db.session.Close()
}