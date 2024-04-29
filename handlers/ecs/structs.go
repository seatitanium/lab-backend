package ecs

import "time"

type StopInstanceRequest struct {
	InstanceId string `json:"instanceId"`
	Force      bool   `json:"force"`
}

type InvocationResult struct {
	Status       string    `json:"status"`
	Content      string    `json:"content"`
	ErrorInfo    string    `json:"errorInfo"`
	StartTime    time.Time `json:"startTime"`
	FinishedTime time.Time `json:"finishedTime"`
}
