package utils

import "strings"

func GetPxeFileNameByMac(mac string) string {
	mac = strings.Replace(mac, ":", "-", -1)
	return "01-" + mac
}
