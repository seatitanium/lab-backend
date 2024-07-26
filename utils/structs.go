package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTPayload struct {
	Username  string `json:"username"`
	UpdatedAt string `json:"updatedAt"`
}

type JWTClaims struct {
	jwt.StandardClaims
	Payload JWTPayload `json:"payload"`
}

type MCIDUsage struct {
	Used     bool   `json:"used"`
	Verified bool   `json:"verified"`
	With     string `json:"with"`
}

type LoginRecordBoard struct {
	Player        string    `json:"player"`
	Count         int       `json:"count"`
	LastCreatedAt time.Time `json:"lastCreatedAt"`
}

type PlaytimeBoard struct {
	Player    string `json:"player"`
	TimeTotal int32  `json:"timeTotal"`
	TimeAfk   int32  `json:"timeAfk"`
}

type TermDownloadItem struct {
	Name string `json:"name,omitempty"`
	Size string `json:"size,omitempty"`
}

type TermDownloads struct {
	World TermDownloadItem `json:"world,omitempty"`
	Pack  TermDownloadItem `json:"pack,omitempty"`
	Mods  TermDownloadItem `json:"mods,omitempty"`
}

type Term struct {
	Tag         string        `json:"tag"`
	Version     string        `json:"version"`
	Theme       string        `json:"theme"`
	ThemeAlt    string        `json:"themeAlt,omitempty"`
	PackVersion string        `json:"packVersion,omitempty"`
	Type        string        `json:"type"`
	Author      string        `json:"author"`
	Link        string        `json:"link,omitempty"`
	StartAt     string        `json:"startAt"`
	EndAt       string        `json:"endAt,omitempty"`
	Created     string        `json:"created"`
	Image       string        `json:"image,omitempty"`
	Downloads   TermDownloads `json:"downloads"`
}

type ServerPlayer struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
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
	Id           uint   `json:"id" gorm:"primaryKey,NOT NULL,AUTO_INCREMENT"`
	Username     string `json:"username" gorm:"size:50;NOT NULL"`
	Nickname     string `json:"nickname" gorm:"size:50;NOT NULL;default:''"`
	Email        string `json:"email" gorm:"size:100;NOT NULL"`
	MCID         string `json:"mcid" gorm:"size:30;NOT NULL;default:''"`
	CreatedAt    int64  `json:"createdAt" gorm:"autoCreateTime:milli;NOT NULL"`
	UpdatedAt    int64  `json:"updatedAt" gorm:"autoUpdateTime:milli;NOT NULL"`
	Hash         string `json:"hash" gorm:"size:512;NOT NULL"`
	MCIDVerified bool   `json:"mcidVerified" gorm:"NOT NULL;default:false"`
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

type LoginRecord struct {
	Id         uint      `json:"id" gorm:"primaryKey,NOT NULL,AUTO_INCREMENT"`
	ActionType bool      `json:"actionType" gorm:"NOT NULL"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime:milli;NOT NULL"`
	Tag        string    `json:"tag" gorm:"size:20;NOT NULL"`
	Player     string    `json:"player" gorm:"size:20;NOT NULL"`
	UUID       string    `json:"uuid" gorm:"size:50;NOT NULL"`
}

type PlaytimeRecord struct {
	Id        uint      `json:"id" gorm:"primaryKey,NOT NULL,AUTO_INCREMENT"`
	Total     int32     `json:"total" gorm:"NOT NULL;default:0"`
	Afk       int32     `json:"afk" gorm:"NOT NULL;default:0"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:milli;NOT NULL"`
	Tag       string    `json:"tag" gorm:"size:20;NOT NULL"`
	Player    string    `json:"player" gorm:"size:20;NOT NULL"`
	UUID      string    `json:"uuid" gorm:"size:50;NOT NULL"`
}

type SnapshotOnlinePlayers struct {
	Id        uint      `json:"id" gorm:"primaryKey,NOT NULL,AUTO_INCREMENT"`
	Count     int       `json:"count" gorm:"NOT NULL;default:0"`
	Names     string    `json:"names" gorm:"NOT NULL;default:''"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:milli;NOT NULL"`
}
