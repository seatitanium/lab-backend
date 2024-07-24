package utils

import (
	errs "errors"
	"gorm.io/gorm"
	"seatimc/backend/errors"
	"sort"
	"time"
)

func GetLoginRecordsByName(playername string, tag string, offset int, limit int) ([]LoginRecord, *errors.CustomErr) {
	conn := GetStatsDBConn()
	var loginRecord []LoginRecord

	result := conn.Where(&LoginRecord{Player: playername, Tag: tag}).Order("created_at desc").Offset(offset).Limit(limit).Find(&loginRecord)
	if result.Error != nil {
		return nil, errors.DbError(result.Error)
	}

	return loginRecord, nil
}

func GetPlaytimeRecord(playername string, tag string) (*PlaytimeRecord, *errors.CustomErr) {
	conn := GetStatsDBConn()
	var playtimeRecord PlaytimeRecord

	if tag == "" {
		tag = GetActiveTerm().Tag
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
		tag = GetActiveTerm().Tag
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

func GetLoginRecordBoard(tag string, limit ...int) ([]LoginRecordBoard, *errors.CustomErr) {
	conn := GetStatsDBConn()
	var loginRecord []LoginRecord

	if tag == "" {
		tag = GetActiveTerm().Tag
	}

	result := conn.Where(&LoginRecord{Tag: tag}).Find(&loginRecord)

	if result.Error != nil {
		return nil, errors.DbError(result.Error)
	}

	var loginRecordBoard = make([]LoginRecordBoard, 0)
	var indexs = make(map[string]int)
	i := 0

	for _, x := range loginRecord {
		if x.ActionType == false {
			continue
		}

		if !HasKey(indexs, x.Player) {
			loginRecordBoard = append(loginRecordBoard, LoginRecordBoard{
				Player:        x.Player,
				Count:         1,
				LastCreatedAt: x.CreatedAt,
			})
			indexs[x.Player] = i
			i += 1
		} else {
			loginRecordBoard[indexs[x.Player]].Count += 1
			loginRecordBoard[indexs[x.Player]].LastCreatedAt = x.CreatedAt
		}
	}

	if len(limit) > 0 {
		return LimitSlice(limit[0], loginRecordBoard), nil
	} else {
		return loginRecordBoard, nil
	}
}

func GetPlaytimeBoard(tag string, limit ...int) ([]PlaytimeBoard, *errors.CustomErr) {
	conn := GetStatsDBConn()
	var playtimeRecord []PlaytimeRecord

	if tag == "" {
		tag = GetActiveTerm().Tag
	}

	result := conn.Where(&LoginRecord{Tag: tag}).Find(&playtimeRecord)

	if result.Error != nil {
		return nil, errors.DbError(result.Error)
	}

	var playtimeRecordBoard = make([]PlaytimeBoard, 0)
	var indexs = make(map[string]int)
	i := 0

	for _, x := range playtimeRecord {
		if !HasKey(indexs, x.Player) {
			playtimeRecordBoard = append(playtimeRecordBoard, PlaytimeBoard{
				Player:    x.Player,
				TimeAfk:   x.Afk,
				TimeTotal: x.Total,
			})
			indexs[x.Player] = i
			i += 1
		}
	}

	sort.Slice(playtimeRecordBoard, func(i, j int) bool {
		return (playtimeRecordBoard[i].TimeTotal - playtimeRecordBoard[i].TimeAfk) > (playtimeRecordBoard[j].TimeTotal - playtimeRecordBoard[j].TimeAfk)
	})

	if len(limit) == 0 || limit[0] > len(playtimeRecordBoard) {
		return playtimeRecordBoard, nil
	} else {
		return playtimeRecordBoard[:limit[0]], nil
	}
}

func GetTermsInvolved(mcid string) ([]Term, *errors.CustomErr) {
	conn := GetStatsDBConn()

	var loginRecords []LoginRecord
	var count int64
	var involved []Term

	for _, t := range GlobalConfig.Terms {
		result := conn.Where(&LoginRecord{Player: mcid, Tag: t.Tag}).Find(&loginRecords).Count(&count)

		if result.Error != nil {
			return nil, errors.DbError(result.Error)
		}

		if count > 0 || HistoryTermsContainsPlayer(t.Tag, mcid) {
			t.StartAt += "T00:00:00Z"
			if t.EndAt != "" {
				t.EndAt += "T00:00:00Z"
			}
			involved = append(involved, t)
		}
	}

	return involved, nil
}

func GetFirstLoginRecord(mcid string, tag string) (*LoginRecord, *errors.CustomErr) {
	conn := GetStatsDBConn()
	var loginRecord LoginRecord

	var query *LoginRecord

	if tag == "" {
		query = &LoginRecord{Player: mcid}
	} else {
		query = &LoginRecord{Player: mcid, Tag: tag}
	}

	result := conn.Where(query).Order("created_at").First(&loginRecord)

	if result.Error != nil {
		if errs.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.DbError(result.Error)
	}

	return &loginRecord, nil
}
