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
	dbutils.CreateSchema()
	usedIps := dbutils.GetUsedIps()
	fmt.Println(usedIps)
	wg.GetWgCIDR()
	config, _ := conf.GetOktaServerConfig()
	oh := okta.OktaHandler{Conf: config}
	oh.GetUsers()
}


