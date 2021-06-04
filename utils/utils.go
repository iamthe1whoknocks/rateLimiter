package utils

import (
	"net"
	"strings"
)

//Получение подсети из ip-адреса, полученного из запроса, и маски подсети.
func GetSubnetFromIP(ip string, mask string) (string, error) {
	var address strings.Builder
	address.WriteString(ip)
	address.WriteString("/")
	address.WriteString(mask)
	_, subnet, err := net.ParseCIDR(address.String())
	if err != nil {
		return "", err
	}
	return subnet.String(), nil
}
