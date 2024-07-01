package utils

import "github.com/golang-jwt/jwt"

type JWTPayload struct {
	Username  string `json:"username"`
	UpdatedAt string `json:"updatedAt"`
}

type JWTClaims struct {
	jwt.StandardClaims
	Payload JWTPayload `json:"payload"`
}

/**
 * Status:
 * 1 - Pending - 创建中
 * 2 - Running - 运行中
 * 3 - Starting - 启动中
 * 4 - Stopping - 停止中
 * 5 - Stopped - 已停止
 */

type Ecs struct {
	Id           uint    `json:"id" gorm:"primaryKey,NOT NULL,AUTO_INCREMENT"`
	InstanceId   string  `json:"instanceId" gorm:"size:50;NOT NULL"`
	TradePrice   float32 `json:"tradePrice" gorm:"NOT NULL"`
	RegionId     string  `json:"regionId" gorm:"size:20;NOT NULL"`
	InstanceType string  `json:"instanceType" gorm:"size:20;NOT NULL"`
	Active       bool    `json:"active" gorm:"NOT NULL;default:true"`
	Status       string  `json:"status" gorm:"size:20;NOT NULL;"`
	DeployStatus string  `json:"deployStatus" gorm:"size:20;default:'Pending'"`
	CreatedAt    int64   `json:"createdAt" gorm:"autoCreateTime:milli;NOT NULL"`
	UpdatedAt    int64   `json:"updatedAt" gorm:"autoUpdateTime:milli;NOT NULL"`
	Ip           string  `json:"ip" gorm:"size:30;default:''"`
}

type Users struct {
	Id        uint   `json:"id" gorm:"primaryKey,NOT NULL,AUTO_INCREMENT"`
	Username  string `json:"username" gorm:"size:50;NOT NULL"`
	Nickname  string `json:"nickname" gorm:"size:50;default:NULL"`
	Email     string `json:"email" gorm:"size:100;NOT NULL"`
	MCID      string `json:"mcid" gorm:"size:30;NOT NULL"`
	CreatedAt int64  `json:"createdAt" gorm:"autoCreateTime:milli;NOT NULL"`
	UpdatedAt int64  `json:"updatedAt" gorm:"autoUpdateTime:milli;NOT NULL"`
	Hash      string `json:"hash" gorm:"size:512;NOT NULL"`
}

type EcsActions struct {
	Id         uint   `json:"id" gorm:"primaryKey,NOT NULL,AUTO_INCREMENT"`
	InstanceId string `json:"instanceId" gorm:"size:50;default:NULL"`
	ActionType string `json:"actionType" gorm:"size:50;NOT NULL"`
	ByUsername string `json:"byUsername" gorm:"size:50;default:NULL"`
	Automated  bool   `json:"automated" gorm:"NOT NULL;default:false"`
	At         int64  `json:"at" gorm:"autoCreateTime:milli;NOT NULL"`
}

type EcsInvokes struct {
	Id         uint   `json:"id" gorm:"primaryKey,NOT NULL,AUTO_INCREMENT"`
	InstanceId string `json:"instanceId" gorm:"size:50;NOT NULL"`
	InvokeId   string `json:"invokeId" gorm:"size:256;NOT NULL"`
	CommandId  string `json:"commandId" gorm:"size:256;NOT NULL"`
	At         int64  `json:"at" gorm:"autoCreateTime:milli;NOT NULL"`
}
