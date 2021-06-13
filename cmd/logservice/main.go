package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http.//%s:%s", host, port)
	r := registry.Registration{
		ServiceName: "Log service",
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(
		context.Background(),
		r,
		host,
		port,

		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()
	fmt.Println("shutting log service.")
}
