/**
 * @author  qyanzh
 * @create  2022/03/08 16:59
 */

package service

import (
	"testing"
)

var schedulerService *SchedulerService

func init() {
	schedulerService = NewSchedulerService()
}

func TestStart(t *testing.T) {
	_ = schedulerService.Start()
	//time.Sleep(10 * time.Second)
	//_ = schedulerService.Stop()
	for {

	}
}
