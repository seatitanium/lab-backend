package monitor

import (
	"log"
	"seatimc/backend/aliyun"
	"seatimc/backend/ecs"
	"seatimc/backend/errors"
	"seatimc/backend/utils"
	"time"
)

// 实例状态监控器
//
// 参数：
//   - interval - 检测时间间隔
//   - threshold - 删除阈值。当检测到停机时间超过该阈值时，执行强制删除。警告：阈值不宜过低（一分钟以上即可），否则可能导致实例在不恰当的时机被删除。
//   - end - 控制监控器的结束状态
func RunInstanceStatusMonitor(interval time.Duration, threshold time.Duration, end <-chan bool) {
	var stoppedDuration time.Duration

	log.Printf("Stopped Instance Monitor\n")
	log.Printf("Running with argument: interval=%v, threshold=%v\n", interval, threshold)

	for {
		var customErr *errors.CustomErr
		var hasActiveInstance bool
		var activeInstance *utils.Ecs
		var retrieved *aliyun.InstanceDescriptionRetrieved

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

		retrieved, customErr = ecs.DescribeInstance(activeInstance.InstanceId, activeInstance.RegionId)

		if customErr != nil {
			log.Println("Critical. Failed DescribeInstance: " + customErr.Handle().Error())
			goto endOfLoop
		}

		if !retrieved.Exist {
			log.Println("Skipped. No result retrieved.")
			goto endOfLoop
		}

		customErr = utils.SetStatus(activeInstance.InstanceId, retrieved.Status)

		if customErr != nil {
			log.Println("Critical. Cannot update instance status: " + customErr.Handle().Error())
			goto endOfLoop
		}

		if retrieved.Status == "Stopped" {
			stoppedDuration += interval
			log.Printf("Retrieved status \"Stopped\", adding stopped duration by %v.\n", interval)
		} else {
			stoppedDuration = 0
			log.Printf("Retrieved status \"%v\", setting stopped duration to 0s.\n", retrieved.Status)
		}

		if stoppedDuration >= threshold {
			log.Printf("Reaching the threshold of inactivity (%v). Forcefully deleting instance %v.\n", threshold, activeInstance.InstanceId)

			customErr = ecs.DeleteInstance(activeInstance.InstanceId, true)
			if customErr != nil {
				log.Println("Critical. Failed DeleteInstance: " + customErr.Handle().Error())
				goto endOfLoop
			}

			customErr = utils.WriteAutomatedEcsRecord(activeInstance.InstanceId, "delete", true)
			if customErr != nil {
				log.Println("Critical. Unable to write automated record: " + customErr.Handle().Error())
				goto endOfLoop
			}

			customErr = utils.SetActive(activeInstance.InstanceId, false)
			if customErr != nil {
				log.Println("Critical. Unable to set instance active state.")
				goto endOfLoop
			}

			log.Println("Finished. Successfully deleted instance.")
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
