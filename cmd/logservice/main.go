package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/sebGilR/dist_services/log"
	"github.com/sebGilR/dist_services/service"
)

func main() {
	log.Run("./dist_services.log")

	host, port := "localhost", "4000"

	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port, log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service")
}
