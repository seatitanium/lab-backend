package monitor

import (
	"log"
	"seatimc/backend/ecs"
	"seatimc/backend/errHandler"
	"seatimc/backend/utils"
	"time"
)

// 实例停止监控器
//
// 注意：必须以 gorountine 执行
//
// 参数：
//   - interval - 检测时间间隔
//   - threshold - 删除阈值。当检测到停机时间超过该阈值时，执行强制删除
//   - end - 控制监控器的结束状态
func RunStoppedInstanceMonitor(interval time.Duration, threshold time.Duration, end <-chan bool) {
	var stoppedDuration time.Duration

	for {
		var customErr *errHandler.CustomErr
		var hasActiveInstance bool
		var activeInstance *utils.Ecs
		var retrieved *ecs.InstanceDescriptionRetrieved

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

		if retrieved.Status == "Stopped" {
			stoppedDuration += interval
		}

		if stoppedDuration >= threshold {
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
