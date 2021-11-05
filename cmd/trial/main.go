/*
wgctrl.New() is the function exposed by the package to create a new client.
 */

package main

import (
	"fmt"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func main()  {
	client,err  := wgctrl.New()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(client) // client is created without any devices
	// next we need to configure devices. Inputs are name and config.
	// to create config, we need to create one or more peer config. 
	cfg := CreateConfig()
	fmt.Println(cfg)
	err = client.ConfigureDevice("wg0",cfg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("no error")
	}
}

func CreateConfig() wgtypes.Config {
	privateKey,_ := wgtypes.GeneratePrivateKey()
	listenPort := 51820
	cfg := wgtypes.Config{
		PrivateKey:   &privateKey,
		ListenPort:   &listenPort,
		FirewallMark: nil,
		ReplacePeers: false,
		Peers:        nil,
	}
	return cfg
}
