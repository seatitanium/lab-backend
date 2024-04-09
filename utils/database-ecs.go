package utils

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"seatimc/backend/ecs"
)

// 使用 instance id 获得数据库中的某个 instance 的记录
func GetInstanceByInstanceId(db *sqlx.DB, instanceId string) (*DbInstance, error) {
	var result DbInstance

	err := DbGet(db, &result, "SELECT * FROM `seati_ecs` WHERE `instance_id`=?", instanceId)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// 获取当前的 active instance
//
// 使用前必须先检验是否存在 active instance，否则会产生 error
func GetActiveInstance(db *sqlx.DB) (*DbInstance, error) {
	var result DbInstance

	err := DbGet(db, &result, "SELECT * FROM `seati_ecs` WHERE `active`=1")

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// 检查当前是否存在 active instance
//
// 注意：第一个返回值为 false 时，不一定表示不存在，也有可能是发生了错误
func HasActiveInstance(db *sqlx.DB) (bool, error) {
	rows, err := DbQuery(db, "SELECT * FROM `seati_ecs` WHERE `active`=1")

	if err != nil {
		return false, err
	}

	hasNext := rows.Next()
	err = rows.Err()

	if !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}

	return hasNext, nil
}

// 将一个 *CreatedInstance 插入数据库，并将其设定为 active
//
// 提醒：插入数据库时，记录的 active 值默认为 true。
func SaveNewActiveInstance(db *sqlx.DB, instance *ecs.CreatedInstance) error {
	_, err := DbExec(db, "INSERT INTO `seati_ecs` (`instance_id`, `trade_price`) VALUES (?, ?)", instance.InstanceId, instance.TradePrice)

	if err != nil {
		return err
	}

	return nil
}
