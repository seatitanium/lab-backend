package ecs

type CommonInstanceRequest struct {
	InstanceId string `json:"instance_id"`
}

type StopInstanceRequest struct {
	InstanceId string `json:"instance_id"`
	Force      bool   `json:"force"`
}
