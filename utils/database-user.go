package utils

import "seatimc/backend/errors"

func IsUsernameUsed(username string) (bool, *errors.CustomErr) {
	conn := GetDBConn()
	var count int64

	result := conn.Model(&Users{}).Where(&Users{Username: username}).Count(&count)

	if result.Error != nil {
		return false, errors.DbError(result.Error)
	}

	return count > 0, nil
}

func GetUserByUsername(username string) (*Users, *errors.CustomErr) {
	conn := GetDBConn()
	var user Users

	result := conn.Where(&Users{Username: username}).Limit(1).Find(&user)

	if result.Error != nil {
		return nil, errors.DbError(result.Error)
	}

	if user.Hash == "" {
		return nil, errors.TargetNotExist()
	}

	return &user, nil
}

// Note: the second argument must be of type map[string]any.
func UpdateUserByUsername(username string, updates map[string]any) *errors.CustomErr {
	conn := GetDBConn()

	result := conn.Model(&Users{}).Where(&Users{Username: username}).Updates(updates)

	if result.Error != nil {
		return errors.DbError(result.Error)
	}

	return nil
}
