package main

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/services", &registry.RegistryService{})
	ctx, cannel := context.WithCancel(context.Background())
	defer cannel()
	var srv http.Server
	srv.Addr = registry.ServerPort
	go func() {
		log.Println(srv.ListenAndServe())
		cannel()
	}()
	go func() {
		fmt.Println("Registry service start")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cannel()
	}()
	<-ctx.Done()
	fmt.Println("shutdown Registry Service")
}
