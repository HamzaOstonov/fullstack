package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
)

//Postgres strukturasi
type Postgres struct {
	db *sqlx.DB
}

//Test strukturasi
type Test struct {
	ID      uint64         `json:"id"`
	Title   string         `json:"title"`
	Answers types.JSONText `json:"answers"`
	Created time.Time      `json:"created,omitempty"`
	Updated time.Time      `json:"updated,omitempty"`
}
