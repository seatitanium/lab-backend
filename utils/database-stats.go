package utils

import "seatimc/backend/errors"

func GetLoginRecordsByName(playername string, tag string, offset int, limit int) ([]LoginRecord, *errors.CustomErr) {
	conn := GetStatsDBConn()
	var loginRecord []LoginRecord

	result := conn.Where(&LoginRecord{Player: playername, Tag: tag}).Offset(offset).Limit(limit).Find(&loginRecord)
	if result.Error != nil {
		return nil, errors.DbError(result.Error)
	}

	return loginRecord, nil
}

func GetPlaytimeRecord(playername string, tag string) (*PlaytimeRecord, *errors.CustomErr) {
	conn := GetStatsDBConn()
	var playtimeRecord PlaytimeRecord

	if tag == "" {
		tag = "default"
	}

	result := conn.Where(&PlaytimeRecord{Player: playername, Tag: tag}).Find(&playtimeRecord)
	if result.Error != nil {
		return nil, errors.DbError(result.Error)
	}

	return &playtimeRecord, nil
}

func GetLoginRecordCount(playername string, tag string) (int64, *errors.CustomErr) {
	conn := GetStatsDBConn()
	var count int64
	var loginRecord []LoginRecord

	if tag == "" {
		tag = "default"
	}

	// Note: Must Find before Count
	result := conn.Where(&LoginRecord{Player: playername, Tag: tag, ActionType: true}).Find(&loginRecord).Count(&count)
	if result.Error != nil {
		return 0, errors.DbError(result.Error)
	}

	return count, nil
}
