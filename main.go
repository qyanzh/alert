package main

import (
	api "alert/kitex_gen/api/combineservice"
	"log"
)

func main() {
	svr := api.NewServer(new(CombineServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
