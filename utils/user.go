package utils

import "slices"

func IsAdmin(username string) bool {
	return slices.Contains(GlobalConfig.Administrators, username)
}
