package utils

import "strings"

// 根据 handlerName 判断该 handler 是否需要验证
func NeedAuthorize(handlerName string) bool {
	needs := GlobalConfig.NeedAuthorizeHandlers
	for i := 0; i < len(needs); i++ {
		if strings.Contains(handlerName, needs[i]) {
			return true
		}
	}
	return false
}
