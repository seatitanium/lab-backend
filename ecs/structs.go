package ecs

type AvailableZone struct {
	// 该可用区内可使用的实例类型
	AvailableInstanceTypes []string `json:"available_instance_types"`
	// 可用区 ID
	ZoneId string `json:"zone_id"`
	// 可用区的本地名称
	ZoneLocalName string `json:"zone_local_name"`
}

type SpotPriceHistory struct {
	// 现价
	Price float32 `json:"price"`
	// 原价
	OriginPrice float32 `json:"origin_price"`
	// 以 ISO8601 / RFC3339 格式表示的时间字符串
	TimeISO8601 string `json:"time_iso_8601"`
	// 时间字符串的时间戳（毫秒）
	Timestamp int64 `json:"time"`
	// 实例类型
	InstanceType string `json:"instance_type"`
	// 可用区 ID
	ZoneId string `json:"zone_id"`
	// 网络类型
	NetworkType string `json:"network_type"`
}

type CreatedInstance struct {
	// 成交价格
	TradePrice float32 `json:"trade_price"`
	// 实例 ID
	InstanceId string `json:"instance_id"`
}
