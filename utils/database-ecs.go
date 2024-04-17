package utils

import "seatimc/backend/errHandler"

// 使用 instance id 获得数据库中的某个 instance 的记录
func GetInstanceByInstanceId(instanceId string) (*Ecs, *errHandler.CustomErr) {
	conn := GetDBConn()
	var ecs Ecs

	result := conn.Where(&Ecs{InstanceId: instanceId}).Limit(1).Find(ecs)
	if result.Error != nil {
		return nil, errHandler.DbError(result.Error)
	}

	return &ecs, nil
}

// 获取当前的 active instance
//
// 使用前必须先检验是否存在 active instance，否则会产生 errHandler
func GetActiveInstance() (*Ecs, *errHandler.CustomErr) {
	conn := GetDBConn()
	var ecs Ecs

	result := conn.Where(&Ecs{Active: true}).Limit(1).Find(ecs)
	if result.Error != nil {
		return nil, errHandler.DbError(result.Error)
	}

	return &ecs, nil
}

// 检查当前是否存在 active instance
//
// 注意：第一个返回值为 false 时，不一定表示不存在，也有可能是发生了错误
func HasActiveInstance() (bool, *errHandler.CustomErr) {
	conn := GetDBConn()
	var ecsCount int64

	result := conn.Where(&Ecs{Active: true}).Count(&ecsCount)
	if result.Error != nil {
		return false, errHandler.DbError(result.Error)
	}

	return ecsCount > 1, nil
}

// 将一个 *CreatedInstance 插入数据库，并将其设定为 active
//
// 提醒：插入数据库时，记录的 active 值默认为 true。
func SaveNewActiveInstance(instance *CreatedInstance, regionId string, instanceType string) *errHandler.CustomErr {
	conn := GetDBConn()

	result := conn.Create(&Ecs{
		InstanceId:   instance.InstanceId,
		TradePrice:   instance.TradePrice,
		RegionId:     regionId,
		InstanceType: instanceType,
	})

	return errHandler.DbError(result.Error)
}
