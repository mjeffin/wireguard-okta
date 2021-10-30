package wg

import (
	"fmt"
	"net"
	"os"
)

func GetWgCIDR()  {
	wgIntIpString := os.Getenv("WG_INTERFACE_IP")
	addr,net,err := net.ParseCIDR(wgIntIpString)
	fmt.Println(addr,net,err)
}