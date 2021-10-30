package dbutils

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	DBFILE = "./wgokta.db"
)

func CreateSchema()  {
	db, err := sql.Open("sqlite3", DBFILE)
	if err != nil {
		log.Println("Error opening DB", err)
	}
	defer db.Close()
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
	}
}

// GetUsedIPs returns the list of ip addresses as strings currently being in use.
// This is called when adding a new user to find the next available ip to be allocated
// They are returned as strings so that the last octat could be easily removed
func GetUsedIPs() []string{
	db, err := sql.Open("sqlite3", DBFILE)
	if err != nil {
		log.Println("Error opening DB", err)
	}
	defer db.Close()
	//https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
	//pgxscan was much more easier!
	rows, err := db.Query("select ip from user")
	if err != nil {
		log.Println("Unable to query from db")
	}
	var ipList []string
	for rows.Next() {
		var i string
		err = rows.Scan(&i)
		if err != nil {
			log.Println("Error scanning result to string")
		}
		ipList = append(ipList, i)
	}
	err = rows.Err()
	if err != nil {
		log.Println("Error fetching ip ", err)
	}
	return ipList
}