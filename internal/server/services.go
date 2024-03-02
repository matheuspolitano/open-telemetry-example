package server

import "github.com/jmoiron/sqlx"

type ServerServices struct {
	db *sqlx.DB
}
