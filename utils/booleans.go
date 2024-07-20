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

// 根据 URI 前缀判断该 URI 请求是否必须来自服务器
func IsServerOnly(uri string) bool {
	needs := GlobalConfig.ServerOnlyEndpoints

	for i := 0; i < len(needs); i++ {
		if strings.HasPrefix(uri, needs[i]) {
			return true
		}
	}

	return false
}

func IsTrue(boolean string) bool {
	return strings.EqualFold(boolean, "true")
}
