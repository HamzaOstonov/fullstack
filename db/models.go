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
	Correct string         `json:"correct,omitempty"`
	Type    uint           `json:"type,omitempty"`
	Created time.Time      `json:"created,omitempty"`
	Updated time.Time      `json:"updated,omitempty"`
}

//User -> Foydalanuvchi strukturasi
type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
