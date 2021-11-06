/*
This program is to try out the wireguard-ctrl package and wg in general. Code will be moved to other places after experimentation.

wgctrl.New() is the function exposed by the package to create a new client. create a client, create a device configuration.
Creating a device configuration actually applies the config the local system!! Tested this in a linux box

 */

package main

import (
	"fmt"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"net"
)

func main()  {
	client,err  := wgctrl.New()
	if err != nil {
		fmt.Println(err)
	}
	// next we need to configure devices. Inputs are name and config.
	// to create config, we need to create one or more peer config. 
	cfg := createConfig()
	err = client.ConfigureDevice("wg0",cfg)
	if err != nil {
		fmt.Println(err)
	}
}


// CreateConfig creates and returns the config to be used go configure the wireguard device
func createConfig() wgtypes.Config {
	privateKey,_ := wgtypes.GeneratePrivateKey()
	listenPort := 51821
	testPeer := createTestPeer()
	cfg := wgtypes.Config{
		PrivateKey:   &privateKey,
		ListenPort:   &listenPort,
		FirewallMark: nil,
		ReplacePeers: false,
		Peers:        []wgtypes.PeerConfig{testPeer},
	}
	return cfg
}

func createTestPeer() wgtypes.PeerConfig {
	ip,netw,_ :=  net.ParseCIDR("10.49.0.2/32")
	mask := netw.Mask
	allowedIps := net.IPNet{
		IP:   ip,
		Mask: mask,
	}
	privateKey,_ := wgtypes.GeneratePrivateKey()
	pubKey := privateKey.PublicKey()
	pc := wgtypes.PeerConfig{
		PublicKey:                   pubKey,
		Remove:                      false,
		UpdateOnly:                  false,
		PresharedKey:                nil,
		Endpoint:                    nil,
		PersistentKeepaliveInterval: nil,
		ReplaceAllowedIPs:           false,
		AllowedIPs:                  []net.IPNet{allowedIps},
	}
	return pc
}