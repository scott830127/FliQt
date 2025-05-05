package main

import (
	"FliQt/internals/app"
	"FliQt/internals/app/config"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfgPath := flag.String("b", "./deploy/config/config.toml", "base config path")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	application, cleanup, err := app.Start(cfg)
	if err != nil {
		log.Fatalf("failed to start app: %v", err)
	}
	defer cleanup()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err = application.Engine.Run(); err != nil {
			log.Fatalf("server run failed: %v", err)
		}
	}()

	<-quit
}
