package monitor

import (
	"log"
	"seatimc/backend/ecs"
	"seatimc/backend/errors"
	"seatimc/backend/utils"
	"time"
)

func RunDeployStatusMonitor(interval time.Duration, end <-chan bool) {

	log.Printf("Deploy Status Monitor\nRunning with argument: interval=%v", interval)

	for {
		var customErr *errors.CustomErr
		var hasActiveInstance bool
		var activeInstance *utils.Ecs
		var invokeId string
		var invocationResult *ecs.InvocationResult

		hasActiveInstance, customErr = utils.HasActiveInstance()

		if customErr != nil {
			log.Println("Critical. Cannot determine active instance existence: " + customErr.Handle().Error())
			goto endOfLoop
		}

		if !hasActiveInstance {
			log.Println("Skipped. No active instance present.")
			goto endOfLoop
		}

		activeInstance, customErr = utils.GetActiveInstance()

		if customErr != nil {
			log.Println("Critical. Cannot get active instance from database: " + customErr.Handle().Error())
			goto endOfLoop
		}

		invokeId, customErr = utils.GetLastInvokeId(activeInstance.InstanceId)

		if customErr != nil {
			log.Println("Critical. Cannot get last invoke id.")
			goto endOfLoop
		}

		if invokeId == "" {
			log.Println("Skipped. No invocation found.")
			goto endOfLoop
		}

		invocationResult, customErr = ecs.DescribeInvocationResults(invokeId)

		if customErr != nil {
			log.Println("Critical. Cannot get invocation result.")
			goto endOfLoop
		}

		customErr = utils.SetDeployStatus(activeInstance.InstanceId, invocationResult.Status)

		if customErr != nil {
			log.Printf("Critical. Cannot set deploy status to %v.\n", invocationResult.Status)
			goto endOfLoop
		}

		if invocationResult.Status == "Success" {
			log.Println("Skipped. Invocation status is now success, waiting for new operations...")
			goto endOfLoop
		}

		log.Printf("Updating deploy status to %v.\n", invocationResult.Status)

	endOfLoop:
		select {
		case <-end:
			break
		default:
			time.Sleep(interval)
			continue
		}

	}
}
