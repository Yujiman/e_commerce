package utils

import (
	"strings"

	"github.com/Yujiman/e_commerce/auth/aggregatorUser/config"
)

func CheckIp(ip string) bool {

	if strings.HasPrefix(ip, "[::1]") || strings.HasPrefix(ip, "127.0.0.1") {
		return true
	}
	ip = strings.Split(ip, ":")[0]

	allowedIps := config.GetAllowedIp()

	for _, allowedIp := range *allowedIps {
		if string(allowedIp[len(allowedIp)-1]) == "*" {
			if isValidOctets(ip, allowedIp) {
				return true
			}
		} else if ip == allowedIp {
			return true
		}
	}

	return false
}

func isValidOctets(ip string, allowedIp string) bool {
	octets := strings.Split(allowedIp, ".")
	ipOctets := strings.Split(ip, ".")

	for i := 0; i < 3; i++ {
		if octets[i] != ipOctets[i] {
			return false
		}
	}

	return true
}
