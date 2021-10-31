package main

import (
	"fmt"
	"log"
	"mjeffn/wireguard-okta/internal"
	"mjeffn/wireguard-okta/pkg/conf"
)

func main() {
	conf.LoadEnvFile() // load .env file
	// get the list of allowed users from okta
	config, err := conf.GetOktaServerConfig()
	if err != nil {
		log.Println("Error fetching okta config")
	}
	oh := internal.OktaHandler{Conf: config}
	userEmails, err := oh.GetAllowedUsers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userEmails)
	// get the list of active users from okta
	if err := internal.CreateDBSchema(); err != nil {
		log.Println("Error creating db schema ", err)
	}
	// compare both lists
	_, wgNet, _ := internal.GetWgCIDR()
	usedIPs, err := internal.GetActiveIPs()
	nextIp, err := internal.GetNextIp(wgNet, usedIPs)
	if err != nil {
		fmt.Println("Error fetching next ip ", err)
	}
	fmt.Println(nextIp)

}
