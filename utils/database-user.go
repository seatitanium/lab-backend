package utils

func GetUserByUsername(username string) (*Users, error) {
	conn := GetDBConn()
	var user Users

	result := conn.Where(&Users{Username: username}).Limit(1).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
