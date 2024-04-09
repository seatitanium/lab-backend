package utils

import "strings"

// 根据 handlerName 判断该 handler 是否需要验证
func NeedAuthorize(handlerName string) bool {
	needs := Conf().NeedAuthorizeHandlers
	for i := 0; i < len(needs); i++ {
		if strings.HasPrefix(handlerName, needs[i]) {
			return true
		}
	}
	return false
}
