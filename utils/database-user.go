package utils

import "seatimc/backend/errHandler"

func GetUserByUsername(username string) (*Users, *errHandler.CustomErr) {
	conn := GetDBConn()
	var user Users

	result := conn.Where(&Users{Username: username}).Limit(1).Find(&user)
	if result.Error != nil {
		return nil, errHandler.DbError(result.Error)
	}
	return &user, nil
}
