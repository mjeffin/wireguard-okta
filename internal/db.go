package internal

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	DBFILE = "./wgokta.db"
)

func CreateDBSchema() error {
	db, err := sql.Open("sqlite3", DBFILE)
	if err != nil {
		log.Println("Error opening DB", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing db connection inside create db schema")
			return
		}
	}(db)
	sqlStmt := `
	create table if not exists user (
	email text not null ,
	ip text not null ,
	unique (email,ip)
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}

// GetActiveIPs returns the list of ip addresses as strings currently being in use.
// This is called when adding a new user to find the next available ip to be allocated
// They are returned as strings so that the last octat could be easily removed
func GetActiveIPs() ([]string, error) {
	db, err := sql.Open("sqlite3", DBFILE)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error opening db connection")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing db connection inside GetActiveIPs")
			return
		}
	}(db)
	//https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
	//pgxscan was much more easier!
	rows, err := db.Query("select ip from user")
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to query from db")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	var ipList []string
	for rows.Next() {
		var i string
		err = rows.Scan(&i)
		if err != nil {
			log.Println("Error scanning row to string - ", err)
		}
		ipList = append(ipList, i)
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error fetching ip ", err)
	}
	return ipList, nil
}

// GetActiveUsers returns the list of email addresses currently configured in db (and hence wireguard config)
func GetActiveUsers() ([]string, error) {
	db, err := sql.Open("sqlite3", DBFILE)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error opening db connection")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing db connection inside GetActiveIPs")
			return
		}
	}(db)
	//https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
	//pgxscan was much more easier!
	rows, err := db.Query("select email from user")
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to query emails from db")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)
	var emails []string
	for rows.Next() {
		var i string
		err = rows.Scan(&i)
		if err != nil {
			log.Println("Error scanning email row to string - ", err)
		}
		emails = append(emails, i)
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error fetching email ", err)
	}
	return emails, nil
}
