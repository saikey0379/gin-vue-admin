package utils

import (
	"fmt"
	"math"
	"net"
)

// GetNetworkAddress calculates the network address based on the IP address and the number of mask bits.
func GetNetworkAddress(ip string, maskBits int) (net.IP, error) {
	ipAddress := net.ParseIP(ip)
	if ipAddress == nil {
		return nil, fmt.Errorf("Invalid IP address: %s", ip)
	}

	mask := net.CIDRMask(maskBits, 32)
	network := ipAddress.Mask(mask)

	return network, nil
}

func IsIPInSubnet(ipAddress string, network string, maskBits int) bool {
	if ipAddress == "" || network == "" || maskBits == 0 {
		return false
	}
	ip := net.ParseIP(ipAddress)
	networkIP := net.ParseIP(network)  // 子网的网络号
	mask := net.CIDRMask(maskBits, 32) // 子网的掩码位

	// 判断IP地址是否属于子网
	if ip.Mask(mask).Equal(networkIP.Mask(mask)) {
		return true
	} else {
		return false
	}
}

// IsSubnetOf 判断一个网段是否属于另一个网段的子网
func IssubMaskBits(maskBits int) bool {
	return 0 < maskBits && maskBits <= 32
}

// IsSubnetOf 判断一个网段是否属于另一个网段的子网
func IsSubnetOf(network string, maskBits int, subNetwork string, subMaskBits int) bool {
	_, networkIPNet, _ := net.ParseCIDR(fmt.Sprintf("%s/%d", network, maskBits))
	subnetIP, _, _ := net.ParseCIDR(fmt.Sprintf("%s/%d", subNetwork, subMaskBits))

	return networkIPNet.Contains(subnetIP)
}

// IsNetworkConflict 判断一个网段是否与另一个网段冲突
func IsNetworkConflict(network string, maskBits int, subNetwork string, subMaskBits int) bool {
	networkIP, networkIPNet, _ := net.ParseCIDR(fmt.Sprintf("%s/%d", network, maskBits))
	subnetIP, subnetIPNet, _ := net.ParseCIDR(fmt.Sprintf("%s/%d", subNetwork, subMaskBits))
	fmt.Println(networkIPNet.Contains(subnetIP))
	fmt.Println(subnetIPNet.Contains(networkIP))

	return networkIPNet.Contains(subnetIP) || subnetIPNet.Contains(networkIP)
}

// GetTotalIPsInSubnet calculates the total number of IP addresses in a subnet based on the mask bits
func GetTotalIPsInSubnet(ip string, maskBits int) int {
	_, subnet, _ := net.ParseCIDR(fmt.Sprintf("%s/%d", ip, maskBits))
	mask := subnet.Mask

	// Calculate the number of available host addresses in the subnet
	hostBits := 32 - maskLen(mask)
	totalIPs := math.Pow(2, float64(hostBits))

	return int(totalIPs) - 2 // Subtract 2 for network and broadcast addresses
}

// maskLen calculates the number of 1-bits in the IP mask
func maskLen(m net.IPMask) int {
	len := 0
	for _, b := range m {
		for b > 0 {
			len++
			b <<= 1
		}
	}
	return len
}
