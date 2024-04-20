package utils

import "seatimc/backend/errHandler"

func IsUsernameUsed(username string) (bool, *errHandler.CustomErr) {
	conn := GetDBConn()
	var exists bool

	result := conn.Model(&Users{}).Select("count(*) > 0").Where(&Users{Username: username}).Find(&exists)

	if result.Error != nil {
		return false, errHandler.DbError(result.Error)
	}

	return exists, nil
}

func GetUserByUsername(username string) (*Users, *errHandler.CustomErr) {
	conn := GetDBConn()
	var user Users

	result := conn.Where(&Users{Username: username}).Limit(1).Find(&user)

	if result.Error != nil {
		return nil, errHandler.DbError(result.Error)
	}
	return &user, nil
}
