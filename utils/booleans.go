package utils

import "strings"

// 根据 URI 前缀判断该 URI 请求是否需要验证
func NeedAuthorize(uri string) bool {
	needs := GlobalConfig.NeedAuthorizeEndpoints

	for i := 0; i < len(needs); i++ {
		if strings.HasPrefix(uri, needs[i]) {
			return true
		}
	}

	return false
}
