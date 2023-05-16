package main

import (
	"api_go/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()

	serverDoneChan := make(chan os.Signal, 1)

	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	srv := server.New(":8080")

	go srv.ListenAndServe()

	log.Println("server started")

	<-serverDoneChan

	srv.Shutdown(ctx)
	log.Println("server turned off")

}
