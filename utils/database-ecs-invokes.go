package utils

import (
	"seatimc/backend/aliyun"
	"seatimc/backend/errHandler"
)

func WriteInvokeRecord(instanceId string, invokeId string) *errHandler.CustomErr {
	conn := GetDBConn()

	result := conn.Create(&EcsInvokes{
		InstanceId: instanceId,
		CommandId:  aliyun.AliyunConfig.DeployCommandId,
		InvokeId:   invokeId,
	})

	if result.Error != nil {
		return errHandler.DbError(result.Error)
	}

	return nil
}

// 判断是否已经在 instanceId 实例上执行部署指令
func HasInvokedOn(instanceId string) bool {
	conn := GetDBConn()
	var invokedCount int64

	result := conn.Model(&EcsInvokes{}).Where(&EcsInvokes{InstanceId: instanceId}).Count(&invokedCount)
	if result.Error != nil {
		return false
	}

	return invokedCount > 0
}
