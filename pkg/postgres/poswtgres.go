package pkgPostgres

import (
	"database/sql"
	"errors"
	"log/slog"
)

const (
	MAX_TRIES = 3
)

type PGHandler struct {
	db *sql.DB
}

func NewPGHandler(dns string) (*PGHandler, func(), error) {
	counter := MAX_TRIES
	for counter > 0 {
		slog.Info("try to connect to db")
		db, err := sql.Open("postgres", dns)
		if err != nil {
			continue
		}
		return &PGHandler{
			db: db,
		}, func() { _ = db.Close() }, nil
	}
	return nil, nil, errors.New("could not connect to db")
}

func (pg *PGHandler) GetDB() *sql.DB {
	return pg.db
}
