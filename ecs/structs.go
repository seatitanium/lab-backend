package ecs

type AvailableZone struct {
	AvailableInstanceTypes []string `json:"available_instance_types"`
	ZoneId                 string   `json:"zone_id"`
	ZoneLocalName          string   `json:"zone_local_name"`
}

type SpotPriceHistory struct {
	Price        float32 `json:"price"`
	OriginPrice  float32 `json:"origin_price"`
	TimeISO8601  string  `json:"time_iso_8601"`
	Timestamp    int64   `json:"time"`
	InstanceType string  `json:"instance_type"`
	ZoneId       string  `json:"zone_id"`
	NetworkType  string  `json:"network_type"`
}
