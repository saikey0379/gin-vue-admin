package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalIPsInSubnet(t *testing.T) {
	var ipAddress = "192.168.0.1"
	var mask = 24
	var result int

	result = GetTotalIPsInSubnet(ipAddress, mask)
	assert.Equal(t, 254, result, fmt.Sprintf("GetTotalIPsInSubnet(%s,%d)", ipAddress, mask))
}

func TestIsSubnetOf(t *testing.T) {
	var network = "192.168.0.1"
	var maskBits = 24
	var subNetwork = "192.168.0.100"
	var subMaskBits = 25
	var result bool

	result = IsSubnetOf(network, maskBits, subNetwork, subMaskBits)
	assert.Equal(t, true, result, fmt.Sprintf("IsSubnetOf(%s/%d, %s/%d)", network, maskBits, subNetwork, subMaskBits))
}

func TestIsNetworkConflict(t *testing.T) {
	var network = "192.168.0.1"
	var maskBits = 24
	var subNetwork = "192.168.1.100"
	var subMaskBits = 23
	var result bool

	result = IsNetworkConflict(network, maskBits, subNetwork, subMaskBits)
	assert.Equal(t, true, result, fmt.Sprintf("IsSubnetOf(%s/%d, %s/%d)", network, maskBits, subNetwork, subMaskBits))
}
