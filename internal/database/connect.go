package database

import (
	conf "backend/internal/config"
	"database/sql"

	_ "modernc.org/sqlite"
)

type Connection struct {
	Conn *sql.DB
}

func Connect() (*Connection, error) {

	c, err := conf.NewConfig()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite3", c.DB)
	if err != nil {
		return nil, err
	}
	return &Connection{Conn: db}, nil
}

func (c *Connection) Close() {
	c.Close()
}
