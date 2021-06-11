package main

import (
	"context"
	"go-scraper/src/cron"
	"go-scraper/src/dao/database"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	filename := "scraper_log.log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("open log file error: %s\n", err)
	}

	// set configuration
	log.SetOutput(f)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		DisableColors:   true,
	})
	log.SetLevel(log.DebugLevel)
}

func main() {
	// prepare for interrupt signal
	quit := make(chan os.Signal, 1)

	// database initialize & connect
	if code, err := database.Initialize(); err != nil {
		log.Fatalf("%d: failed connect to database, %s\n", code, err)
	}

	// run cron
	cron.Start()

	// wait until shutdown
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// graceful shutdown
	SafeShutdown()
}

func SafeShutdown() {
	// close database
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	database.GetInstance().Destroy()

	// stop cron
	cron.Stop()
}
