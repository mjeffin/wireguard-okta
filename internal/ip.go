package internal

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	ALLOWEDMASK=24
)

// GetWgCIDR reads the ip address for the wg interface in the server and the CIDR network to be used for peers
// It panics when the network mask is not /24 . Extra steps are needed to compute the next available address in that scenario.
// This restriction will be removed in the future once all all other steps are complete.
// TODO - support masks other than /24
func GetWgCIDR() (net.IP, *net.IPNet, error) {
	wgIntIpString := os.Getenv("WG_INTERFACE_IP")
	addr,ipnet,err := net.ParseCIDR(wgIntIpString)
	mask,_ := ipnet.Mask.Size()
	if mask != ALLOWEDMASK {
		log.Printf("Network masks other than %d is not supported at the moment. Exiting",ALLOWEDMASK)
		os.Exit(1)
	}
	return addr,ipnet,err
}

// getLastOctetList returns the list of last octets as integers
// Checks for error while converting to int and also if the number is less than 254
func getLastOctetList(ipList []string) []int {
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

//GetNextIp returns the next available ip that can be used for a new peer or error when it's not possible to assign an ip
// Assumption - the subnet mask is /24.
func GetNextIp(wgNet *net.IPNet,  usedIPs[]string) (net.IP, error) {
	lastOctets := getLastOctetList(usedIPs)
	mask,_ := wgNet.Mask.Size()
	if mask != ALLOWEDMASK {
		log.Printf("Network masks other than %d is not supported at the moment. Exiting",ALLOWEDMASK)
		os.Exit(1)
	}
	sort.Slice(lastOctets, func(i, j int) bool {
		return i<j
	})
	if len(lastOctets) >= getMaxAllowedIps(mask) {
		err := errors.New("reached maximum number of users")
		return net.IP{},err
	}
	var nextOctet int
	for i,j := range lastOctets {
		if i+2 != j {
			nextOctet = i+2 // ip ranges starts with 2 and i starts at 0
			break
		}
		nextOctet = i+3
	}
	log.Printf("Found next last octet - %d. Need to form ip out of it",nextOctet)
	nextIp,err := formIp(wgNet,nextOctet)
	if err != nil {
		return nil,err
	}
	fmt.Println(nextIp)
	return nextIp,nil
}

//getMaxAllowedIps returns the max number of allowed ip addresses(peers) for the given subnet mask
func getMaxAllowedIps(mask int) int  {
	return int(math.Pow(2,float64(mask))) - 3
}

//formIp returns the ip address from the network and last octet.
// works only for /24 mask currently
func formIp(ipNet *net.IPNet, lastOctet int ) (net.IP,error) {
	mask,_ := ipNet.Mask.Size()
	if mask != ALLOWEDMASK {
		log.Printf("Network masks other than %d is not supported at the moment. Exiting",ALLOWEDMASK)
		os.Exit(1)
	}
	nwString := ipNet.IP.String()
	fol := strings.Split(nwString,".")
	firstOctets := fol[:3]
	lastOctetS := strconv.Itoa(lastOctet)
	ipSlice := append(firstOctets,lastOctetS)
	ipString := strings.Join(ipSlice,".")
	ip := net.ParseIP(ipString)
	if !ipNet.Contains(ip) {
		return net.IP{}, errors.New("invalid ip. ip doesn't belong to the wg network")
	}
	return ip,nil
}