package main

import (
	"bot/config"
	"bot/server"
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	cfg, err := config.Load("")
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("recover when start project %v", err)
		}
	}()
	if err != nil {
		log.Fatalf("cannot get config : %v", err)
		panic(err)
	}
	s, err := server.NewServer(cfg)
	if err != nil {
		panic(err)
	}
	s.Init()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	go func() {
		if err := s.StartServer(); err != nil {
			panic(err)
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.HttpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

}
