package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sebGilR/dist_services/registry"
)

func main() {
	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registry started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)

		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Registry stopped.")
}
