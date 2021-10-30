package wg

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// GetWgCIDR reads the ip address for the wg interface in the server and the CIDR network to be used for peers
// It panics when the network mask is less than /24 as extra steps are needed to compute the next available address in that scenario.
// This restriction will be removed in future.
func GetWgCIDR() (net.IP, *net.IPNet, error) {
	wgIntIpString := os.Getenv("WG_INTERFACE_IP")
	addr,ipnet,err := net.ParseCIDR(wgIntIpString)
	log.Println(addr,ipnet,err)
	mask, err := strconv.Atoi(ipnet.Mask.String())
	if err != nil {
		log.Println("Unable to fetch network mask from config")
	}
	if mask < 24 {
		log.Fatal("Network masks less than /24 is not supported at the moment. Exiting")
	}
	return addr,ipnet,err
}

// GetLastOctetList returns the list of last octets as integers
// Checks for error while converting to int and also if the number is less than 254
func GetLastOctetList(ipList []string) []int {
	var octets []int
	for _,ip := range ipList {
		splitted := strings.Split(ip,".")
		if len(splitted) != 4 {
			continue
		}
		octetInt,err := strconv.Atoi(splitted[3])
		if err != nil || octetInt > 254 {
			log.Println("Unable to convert the last octet to int")
		}
		octets = append(octets, octetInt)
	}
	return octets
}

// GetNextIp returns the next available ip that can be used for a new peer or error when it's not possible to assign an ip
// Assumption - the
//func GetNextIp(wgIntIp net.IP, wgNet *net.IPNet, usedIPs []net.IP) (net.IP,error) {
//
//}