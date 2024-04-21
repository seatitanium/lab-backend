package ecs

type StopInstanceRequest struct {
	InstanceId string `json:"instanceId"`
	Force      bool   `json:"force"`
}
