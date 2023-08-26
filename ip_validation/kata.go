package kata

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	octets := strings.Split(ip, ".")
	if len(octets) != 4 {
		return false
	}
	for _, octet := range octets {
		if len(octet) > 1 && octet[0] == '0' {
			return false
		}
		v, err := strconv.Atoi(octet)
		if err != nil {
			return false
		}
		if v < 0 || v > 255 {
			return false
		}
	}
	return true
}
