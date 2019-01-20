package main

import (
	"github.com/huyujie/gin-demo/logger"
	"github.com/huyujie/gin-demo/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logger.Setup()

	router := routers.InitRouter()
	// Listen and Server in 0.0.0.0:8080
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		os.Exit(0)
	}()
	select {}
}
