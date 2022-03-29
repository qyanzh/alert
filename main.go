package main

import (
	"alert/internal/service"
	"alert/kitex_gen/api/combineservice"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"log"
)

var schedulerService *service.SchedulerService

func init() {
	schedulerService = service.NewSchedulerService()
}
func main() {
	if err := schedulerService.Start(); err != nil {
		log.Panicln(err)
	}

	addr := utils.NewNetAddr("tcp", ":9999")
	svr := combineservice.NewServer(NewCombineServiceImpl(), server.WithServiceAddr(addr))
	if err := svr.Run(); err != nil {
		log.Panicln(err)
	}
}
