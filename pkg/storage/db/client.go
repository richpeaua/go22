package db

import (
	"database/sql"
	"strconv"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/richpeaua/go22/pkg/adding"
)

type DB struct {
	*sql.DB
}

// Create Sqlite3 DB
func NewDB(repoName string) (*DB, error) {
	db, err := sql.Open("sqlite3", repoName)
	if err != nil {
		return, nil, err
	}
	return &DB{db}, nil
}

func (db *DB) AddConnection()

