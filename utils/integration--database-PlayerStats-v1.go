package utils

import (
	"seatimc/backend/errors"
	"time"
)

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

func GetOnlineHistory(startTime time.Time, endTime time.Time) ([]SnapshotOnlinePlayers, *errors.CustomErr) {
	conn := GetStatsDBConn()
	var snapshots = make([]SnapshotOnlinePlayers, 0)

	result := conn.Where("created_at BETWEEN ? AND ?", startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05")).Find(&snapshots)
	if result.Error != nil {
		return nil, errors.DbError(result.Error)
	}

	return snapshots, nil
}

// 返回最新且 Count 最大的那个 SnapshotOnlinePlayers 记录值
func GetPeakOnlineHistory() (*SnapshotOnlinePlayers, *errors.CustomErr) {
	conn := GetStatsDBConn()
	var snapshots = make([]SnapshotOnlinePlayers, 0)

	result := conn.Find(&snapshots)
	if result.Error != nil {
		return nil, errors.DbError(result.Error)
	}

	maximumIndex := 0
	maximumCount := 0

	for i, snapshot := range snapshots {
		if snapshot.Count > maximumCount {
			maximumIndex = i
		}
	}

	return &snapshots[maximumIndex], nil
}

func GetLoginHistory(startTime time.Time, endTime time.Time) ([]LoginRecord, *errors.CustomErr) {
	conn := GetStatsDBConn()
	var loginRecord = make([]LoginRecord, 0)

	result := conn.Where("created_at BETWEEN ? AND ?", startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05")).Find(&loginRecord)
	if result.Error != nil {
		return nil, errors.DbError(result.Error)
	}

	return loginRecord, nil
}
