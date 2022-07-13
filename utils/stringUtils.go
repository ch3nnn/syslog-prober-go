package utils

import (
	"strings"
)

// IpByColon ip 端口号 通过冒号分割获取 ip
func IpByColon(s string) string {
	split := strings.Split(s, ":")
	return split[0]
}
