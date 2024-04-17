package monitor

import (
	"log"
	"seatimc/backend/ecs"
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
		var err error
		var hasActiveInstance bool
		var activeInstance *utils.Ecs
		var retrieved *ecs.InstanceDescriptionRetrieved

		hasActiveInstance, err = utils.HasActiveInstance()

		if err != nil {
			log.Println("Critical. Cannot determine active instance existence: " + err.Error())
			goto endOfLoop
		}

		if !hasActiveInstance {
			log.Println("Skipped. No active instance present.")
			goto endOfLoop
		}

		activeInstance, err = utils.GetActiveInstance()

		if err != nil {
			log.Println("Critical. Cannot get active instance from database: " + err.Error())
			goto endOfLoop
		}

		retrieved, err = ecs.DescribeInstance(activeInstance.InstanceId, activeInstance.RegionId)

		if err != nil {
			log.Println("Critical. Failed DescribeInstance: " + err.Error())
			goto endOfLoop
		}

		if retrieved.Status == "Stopped" {
			stoppedDuration += interval
		}

		if stoppedDuration >= threshold {
			err = ecs.DeleteInstance(activeInstance.InstanceId, true)
			if err != nil {
				log.Println("Critical. Failed DeleteInstance: " + err.Error())
				goto endOfLoop
			}

			err = utils.WriteAutomatedEcsRecord(activeInstance.InstanceId, "delete", true)

			if err != nil {
				log.Println("Critical. Unable to write automated record: " + err.Error())
				goto endOfLoop
			}
		}

	endOfLoop:
		if <-end {
			break
		} else {
			time.Sleep(interval)
			continue
		}
	}
}
