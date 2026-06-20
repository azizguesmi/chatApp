package database

import (
	conf "backend/internal/config"
	"database/sql"
	"fmt"

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
	fmt.Printf("path %s\n", c.DB)
	db, err := sql.Open("sqlite", c.DB)
	if err != nil {
		return nil, err
	}
	return &Connection{Conn: db}, nil
}

func (c *Connection) Close() {
	c.Conn.Close()
}
