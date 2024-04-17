package ecs

type CommonInstanceRequest struct {
	InstanceId string `json:"instanceId"`
}

type StopInstanceRequest struct {
	InstanceId string `json:"instanceId"`
	Force      bool   `json:"force"`
}
