package utils

import (
	"seatimc/backend/errors"
	"slices"
)

func IsAdmin(username string) bool {
	return slices.Contains(GlobalConfig.Administrators, username)
}

// 从 username 和 playername 两个参数中获得最终的玩家名的值
func GetPlayernameByDoubleProvision(username string, playername string) (string, *errors.CustomErr) {
	if playername != "" {
		return playername, nil
	}

	user, customErr := GetUserByUsername(username)

	if customErr != nil {
		return "", nil
	}

	return user.MCID, nil
}
