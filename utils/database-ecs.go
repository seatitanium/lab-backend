package utils

import (
	"seatimc/backend/aliyun"
	"seatimc/backend/errHandler"
)

// 使用 instance id 获得数据库中的某个 instance 的记录
func GetInstanceByInstanceId(instanceId string) (*Ecs, *errHandler.CustomErr) {
	conn := GetDBConn()
	var ecs Ecs

	result := conn.Where(&Ecs{InstanceId: instanceId}).Limit(1).Find(&ecs)
	if result.Error != nil {
		return nil, errHandler.DbError(result.Error)
	}

	return &ecs, nil
}

// 获取当前的 active instance
//
// 使用前必须先检验是否存在 active instance，否则会产生 customErr
func GetActiveInstance() (*Ecs, *errHandler.CustomErr) {
	conn := GetDBConn()
	var ecs Ecs

	result := conn.Where(&Ecs{Active: true}).Limit(1).Find(&ecs)
	if result.Error != nil {
		return nil, errHandler.DbError(result.Error)
	}

	return &ecs, nil
}

// 检查当前是否存在 active instance
func HasActiveInstance() (bool, *errHandler.CustomErr) {
	conn := GetDBConn()
	var ecsCount int64

	// fixed 04.20:
	// [Database 1302] Msg: [unsupported data type: 0x1400036013e: Table not set, please set it like: db.Model(&user) or db.Table("users")]
	result := conn.Model(&Ecs{}).Where(&Ecs{Active: true}).Count(&ecsCount)
	if result.Error != nil {
		return false, errHandler.DbError(result.Error)
	}

	return ecsCount > 0, nil
}

func GetActiveInstanceId() (string, *errHandler.CustomErr) {
	hasActiveInstance, customErr := HasActiveInstance()

	if customErr != nil {
		return "", customErr
	}

	if !hasActiveInstance {
		return "", errHandler.NotFound()
	}

	activeInstance, customErr := GetActiveInstance()

	if customErr != nil {
		return "", customErr
	}

	return activeInstance.InstanceId, nil
}

// 将一个 *aliyun.CreatedInstance 插入数据库，并将其设定为 active
//
// 提醒：插入数据库时，记录的 active 值默认为 true。
func SaveNewActiveInstance(instance *aliyun.CreatedInstance, regionId string, instanceType string) *errHandler.CustomErr {
	conn := GetDBConn()

	result := conn.Create(&Ecs{
		InstanceId:   instance.InstanceId,
		TradePrice:   instance.TradePrice,
		RegionId:     regionId,
		InstanceType: instanceType,
		Status:       "Pending",
	})

	if result.Error != nil {
		return errHandler.DbError(result.Error)
	}

	return nil
}

func SetStatus(instanceId string, status string) *errHandler.CustomErr {
	conn := GetDBConn()

	result := conn.Model(&Ecs{}).Where(&Ecs{InstanceId: instanceId}).Updates(&Ecs{Status: status})

	if result.Error != nil {
		return errHandler.DbError(result.Error)
	}

	return nil
}

func SetActive(instanceId string, active bool) *errHandler.CustomErr {
	conn := GetDBConn()

	// Note: Must use map[string]any instead of struct itself here.
	// Reference: "When update with struct, GORM will only update non-zero fields, you might want to use map to update attributes or use Select to specify fields to update"
	// https://stackoverflow.com/questions/56653423/gorm-doesnt-update-boolean-field-to-false
	result := conn.Model(&Ecs{}).Where(&Ecs{InstanceId: instanceId}).Updates(map[string]any{"active": active})

	if result.Error != nil {
		return errHandler.DbError(result.Error)
	}

	return nil
}
