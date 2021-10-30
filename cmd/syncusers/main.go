package main

import (
	"fmt"
	"log"
	"mjeffn/wireguard-okta/internal"
	"mjeffn/wireguard-okta/pkg/conf"
)

func main()  {
	conf.LoadEnvFile()
	//_,wgNet,_ := internal.GetWgCIDR()
	//if err := internal.CreateDBSchema(); err != nil {
	//	log.Println("Error creating db schema ", err)
	//}
	//usedIPs := internal.GetUsedIPs()
	//nextIp, err := internal.GetNextIp(wgNet,usedIPs)
	//if err != nil {
	//	fmt.Println("Error fetching next ip ",err)
	//}
	//fmt.Println(nextIp)

	config, err := conf.GetOktaServerConfig()
	if err != nil {
		log.Println("Error fetching okta config")
	}
	oh := internal.OktaHandler{Conf: config}
	userEmails,err := oh.GetAllowedEmails()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userEmails)

}


