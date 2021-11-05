package internal

import (
	"fmt"
	//"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"log"
	"net"
	"time"
)





//CreatePeer creates a peer
func CreatePeer(peerIp net.IPNet) () {
	keyPair, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		log.Println("Error generating key pair")
	}
	psk, err := wgtypes.GenerateKey()
	if err != nil {
		log.Println("Error generating pre shared key")
	}

	 peer := wgtypes.Peer{
		PublicKey:                   keyPair,
		PresharedKey:                psk,
		Endpoint:                    nil,
		PersistentKeepaliveInterval: 0,
		LastHandshakeTime:           time.Time{},
		ReceiveBytes:                0,
		TransmitBytes:               0,
		AllowedIPs:                  []net.IPNet{peerIp},
		ProtocolVersion:             0,
	}
	 fmt.Println(peer)
	 peerConfig := wgtypes.PeerConfig{
		 PublicKey:                   keyPair.PublicKey(),
		 Remove:                      false,
		 UpdateOnly:                  false,
		 PresharedKey:                &psk,
		 Endpoint:                    nil,
		 PersistentKeepaliveInterval: nil,
		 ReplaceAllowedIPs:           false,
		 AllowedIPs:                  nil,
	 }
	 fmt.Println(peerConfig)
}


// Generates WireGuard keys and returns them as strings.
func genKeys() (privateKey string, publicKey string, presharedKey string) {
	keyPair, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		log.Println("Error generating key pair")
	}
	psk, err := wgtypes.GenerateKey()
	if err != nil {
		log.Println("Error generating pre shared key")
	}
	privateKey = keyPair.String()
	publicKey = keyPair.PublicKey().String()
	presharedKey = psk.String()
	return privateKey, publicKey, presharedKey
}

