package backend

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	// "fmt"
)

// Adder interface for any object that can save a connection
type Adder interface {
	AddConn(Connection) error
}

// DB database representation
type DB struct {
	*sql.DB
}

const ( 
	
	connTable  string = `
CREATE TABLE IF NOT EXISTS
connections (id INTEGER PRIMARY KEY, conn_name TEXT NOT NULL, hostname TEXT, 
	ipaddress TEXT, authtype TEXT NOT NULL, username TEXT, password TEXT, privkey TEXT, pubkey TEXT)`

	connInsert string = `
INSERT INTO connections (conn_name, hostname, ipaddress, authtype, username, password, privkey, pubkey) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

)

//NewDB initialize SSH connection database
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	db.Exec(connTable)
	return &DB{db}, nil 
}

//AddConn add SSH connection
func (db *DB) AddConn(conn Connection) error {
	stmt, err := db.Prepare(connInsert)
	if err != nil {
		return err
	}
	defer db.Close()
	stmt.Exec(conn.ConnName, conn.HostName, conn.IPAddress, conn.AuthType, conn.PrivKey, conn.PubKey, conn.Username, conn.Password)
	// defer stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

