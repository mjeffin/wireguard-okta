package main

import (
	"fmt"
	"mjeffn/wireguard-okta/pkg/conf"
	"mjeffn/wireguard-okta/pkg/dbutils"
	"mjeffn/wireguard-okta/pkg/okta"
	"mjeffn/wireguard-okta/pkg/wg"
)

func main()  {
	conf.LoadEnvFile()
	//wg.GetWgCIDR()
	dbutils.CreateSchema()
	usedIPs := dbutils.GetUsedIPs()
	fmt.Println(usedIPs)
	lastOctets := wg.GetLastOctetList(usedIPs)
	fmt.Println(lastOctets)
	config, _ := conf.GetOktaServerConfig()
	oh := okta.OktaHandler{Conf: config}
	oh.GetUsers()
}


