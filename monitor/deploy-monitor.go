package monitor

import (
	"log"
	"seatimc/backend/aliyun"
	"seatimc/backend/ecs"
	"seatimc/backend/errors"
	"seatimc/backend/utils"
	"time"
)

func RunDeployMonitor(interval time.Duration, end <-chan bool) {
	log.Println("Deploy Monitor")
	log.Printf("Running with argument: interval=%v", interval)

	for {
		var hasActiveInstance bool
		var customErr *errors.CustomErr
		var activeInstance *utils.Ecs
		var retrieved *aliyun.InstanceDescriptionRetrieved

		hasActiveInstance, customErr = utils.HasActiveInstance()

		if customErr != nil {
			log.Println("Critical. Error occurred getting active instance:" + customErr.Handle().Error())
			goto endOfLoop
		}

		if hasActiveInstance == false {
			log.Println("Skipped. No active instance present.")
			goto endOfLoop
		}

		activeInstance, customErr = utils.GetActiveInstance()

		if customErr != nil {
			log.Println("Critical. Error occurred getting active instance:" + customErr.Handle().Error())
			goto endOfLoop
		}

		if utils.HasInvokedOn(activeInstance.InstanceId) {
			goto endOfLoop
		}

		retrieved, customErr = ecs.DescribeInstance(activeInstance.InstanceId, activeInstance.RegionId)

		if customErr != nil {
			log.Println("Critical. Failed DescribeInstance: " + customErr.Handle().Error())
			goto endOfLoop
		}

		if !retrieved.Exist {
			log.Println("Skipped. No result retrieved.")
			goto endOfLoop
		}

		if retrieved.Status == "Pending" {
			log.Println("Waiting for instance to be ready...")
			goto endOfLoop
		}

		if retrieved.Status == "Stopped" {
			log.Println("Starting instance.")
			customErr = ecs.StartInstance(activeInstance.InstanceId)

			if customErr != nil {
				log.Println("Critical. Cannot start instance: " + customErr.Handle().Error())
				goto endOfLoop
			}
		}

		if retrieved.Status == "Stopping" || retrieved.Status == "Starting" {
			log.Printf("Warn: The instance is %v but not deployed. Please check if this situation is normal.\n", activeInstance.Status)
			goto endOfLoop
		}

		if retrieved.Status == "Running" {
			log.Println("Detected Running status.")
			log.Println("Starting step 1: IP Address Allocation")

			ip, customErr := ecs.AllocatePublicIpAddress(activeInstance.InstanceId)

			if customErr != nil {
				log.Println(customErr.Handle().Error())
				log.Println("Critical. Cannot allocate ip address. Please try again manually. Skipped.")
			} else {
				log.Printf("Successfully allocated ip address %v for instance %v.", ip, activeInstance.InstanceId)
				customErr = utils.SetInstanceIp(activeInstance.InstanceId, ip)
				if customErr != nil {
					log.Println("Critical. Cannot save IP address to database.")
					log.Println(customErr.Handle().Error())
				} else {
					log.Println("Saved IP address to database.")
				}
			}

			log.Println("Starting step 2: Service Deployment.")
			log.Println("Checking cloud assistant status...")

			var assistantStatusTried = 0

			for {
				assistantStatusTried++
				log.Printf("Trying to get cloud assistant status (%v/%v)\n", assistantStatusTried, 10)

				assistantStatus, customErr := ecs.DescribeCloudAssistantStatus(activeInstance.InstanceId)

				if customErr != nil {
					log.Println("Error occurred getting cloud assistant status:" + customErr.Handle().Error())
				}

				if customErr == nil && assistantStatus == false {
					log.Println("Cloud assistant is not ready. Retrying.")
				}

				if customErr == nil && assistantStatus == true {
					log.Println("Cloud assistant is ready.")
					break
				}

				time.Sleep(time.Second * 5)

				if assistantStatusTried >= 10 {
					log.Println("Reaching maximum trying times. Ending current deployment attempt.")
					goto endOfLoop
				}
			}

			customErr = ecs.InvokeCommand(activeInstance.InstanceId)

			if customErr != nil {
				log.Println("Error occurred trying to invoke command on instance " + activeInstance.InstanceId + ":" + customErr.Handle().Error())
				goto endOfLoop
			}

			log.Println("Successfully invoked command on instance " + activeInstance.InstanceId + ".")
		}

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
