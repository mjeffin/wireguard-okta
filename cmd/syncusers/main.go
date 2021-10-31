package main

import (
	"fmt"
	"log"
	"mjeffn/wireguard-okta/internal"
	"mjeffn/wireguard-okta/pkg/conf"
)

func main() {
	conf.LoadEnvFile()                              // load .env file
	allowedUsers, err := internal.GetAllowedUsers() // from okta
	if err != nil {
		fmt.Println(err)
	}
	// create the db if it's a new deployment
	if err := internal.CreateDBSchema(); err != nil {
		log.Println("Error creating db schema ", err)
	}
	activeUsers, err := internal.GetActiveUsers() // from db. db and wg conf are in sync
	if err != nil {
		log.Println(err)
	}
	usersToAdd, usersToRemove := internal.CompareUsers(allowedUsers, activeUsers)
	log.Printf("Users to add - %s\n", usersToAdd)
	log.Printf("Users to remove - %s\n", usersToRemove)
	//_, wgNet, _ := internal.GetWgCIDR()
	//usedIPs, err := internal.GetActiveIPs()
	//nextIp, err := internal.GetNextIp(wgNet, usedIPs)
	//if err != nil {
	//	log.Println("Error fetching next ip ", err)
	//}
	//log.Println(nextIp)

}
