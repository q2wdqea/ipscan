package main

import (
	"net"
	"strconv"
	"strings"
)

// sip split part to array
func sip(part string) []string {
	parts := strings.Split(part, "-")

	// Convert IPs to uint32 for easy incrementation
	startIPUint := ipToUint32(net.ParseIP(parts[0]))
	endIPUint := ipToUint32(net.ParseIP(parts[1]))

	// Check if start IP is greater than end IP
	if startIPUint > endIPUint {
		return nil
	}

	// Generate the list of IPs
	var ipList []string
	for ip := startIPUint; ip <= endIPUint; ip++ {
		ipList = append(ipList, uint32ToIP(ip).String())
	}

	return ipList
}

// ipToUint32 converts a IPv4 address to an uint32
func ipToUint32(ip net.IP) uint32 {
	bits := strings.Split(ip.String(), ".")

	i0, _ := strconv.Atoi(bits[0])
	i1, _ := strconv.Atoi(bits[1])
	i2, _ := strconv.Atoi(bits[2])
	i3, _ := strconv.Atoi(bits[3])
	return uint32(i0)<<24 | uint32(i1)<<16 | uint32(i2)<<8 | uint32(i3)
}

// uint32ToIP converts a uint32 to an IPv4 address
func uint32ToIP(ipUint uint32) net.IP {
	return net.IPv4(
		byte(ipUint>>24),
		byte(ipUint>>16)&0xFF,
		byte(ipUint>>8)&0xFF,
		byte(ipUint)&0xFF,
	)
}
