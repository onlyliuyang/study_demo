package main

import (
	"context"
	"github.com/testProject/liuyang/gin-example/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := routers.InitRouter()
	routers.InitMetrics(router)

	s := &http.Server{
		Addr:           "127.0.0.1:8090",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
