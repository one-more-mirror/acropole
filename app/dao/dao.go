package dao

import "gitlab.com/one-more/acropole/app/config"

type Dao struct {
	db *Db
}

func NewDao(config config.MongoConfig) (*Dao, error) {
	db, err := newDb(config)

	if err != nil {
		return nil, err
	}

	return &Dao{db: db}, nil
}

func (dao *Dao) Close() {
	dao.db.Close()
}
