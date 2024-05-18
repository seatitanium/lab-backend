package utils

import "seatimc/backend/errHandler"

func IsUsernameUsed(username string) (bool, *errHandler.CustomErr) {
	conn := GetDBConn()
	var count int64

	result := conn.Model(&Users{}).Where(&Users{Username: username}).Count(&count)

	if result.Error != nil {
		return false, errHandler.DbError(result.Error)
	}

	return count > 0, nil
}

func GetUserByUsername(username string) (*Users, *errHandler.CustomErr) {
	conn := GetDBConn()
	var user Users

	result := conn.Where(&Users{Username: username}).Limit(1).Find(&user)

	if result.Error != nil {
		return nil, errHandler.DbError(result.Error)
	}

	if user.Hash == "" {
		return nil, errHandler.TargetNotExist()
	}

	return &user, nil
}
