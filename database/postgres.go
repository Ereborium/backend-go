package database

import (
	"context"

	"github.com/go-pg/pg/v10"
)

type Postgres struct {
	db *pg.DB
}

func NewPostgres(cfg *Config) (Database, error) {
	db := pg.Connect(&pg.Options{
		Addr:        cfg.Address,
		User:        cfg.Username,
		Password:    cfg.Password,
		Database:    cfg.Database,
		DialTimeout: cfg.DialTimeout,
	})

	err := db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (p *Postgres) Close() error {
	return p.db.Close()
}
