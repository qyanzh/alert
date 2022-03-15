package main

import "alert/internal/service"

var schedulerService *service.SchedulerService

func init() {
	schedulerService = service.NewSchedulerService()
}
func main() {
	_ = schedulerService.Start()
	//time.Sleep(10 * time.Second)
	//_ = schedulerService.Stop()
	for {

	}
}
