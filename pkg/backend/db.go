package backend

import (
	"errors"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)


// DB database representation
type DB struct {
	*sql.DB
}

const ( 
	
	connTable  string = `
CREATE TABLE IF NOT EXISTS
connections (id INTEGER PRIMARY KEY, conn_name TEXT UNIQUE NOT NULL, hostname TEXT, 
	ipaddress TEXT, authtype TEXT NOT NULL, username TEXT, password TEXT, privkey TEXT, pubkey TEXT)`

	connInsert string = `
INSERT INTO connections (conn_name, hostname, ipaddress, authtype, username, password, privkey, pubkey) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

)

// NewDB initialize SSH connection database.
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	err = db.Ping()
	if err != nil {
		return nil, errors.New(`Unable to open database: destination directory "data" does not exist
Please check your "$HOME/.go22/" directory or run "go22 init" to reinitalize the application.`)
	}

	_, err = db.Exec(connTable)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil 
}

// AddConn add SSH connection.
func (db *DB) AddConn(conn Connection) error {
	
	// Preparing SQL Insert statement to save connections.
	stmt, err := db.Prepare(connInsert)
	if err != nil {
		return err
	}

	defer db.Close()

	// Executing above SQL Insert with connection attributes passed in.
	_, err = stmt.Exec(conn.ConnName, conn.HostName, conn.IPAddress, conn.AuthType, conn.Username, conn.Password, conn.PrivKey, conn.PubKey)
	if err != nil {
		insertError := strings.Contains(err.Error(), "UNIQUE constraint failed")
		err = errors.New("connection \"" + conn.ConnName +  "\" already exists")
		if insertError {
			return err
		}
	}
	
	return nil
}


