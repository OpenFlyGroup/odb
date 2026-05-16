package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/OpenFlyGroup/odb/internal/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	srv, err := server.New(server.Config{
		DataDir:    "data",
		ListenAddr: ":7200",
	})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		log.Println("shutting down...")
		if err := srv.Shutdown(); err != nil {
			log.Fatalf("shutdown error: %v", err)
		}
	}()

	log.Println("ObjectDB starting on :7200")
	if err := srv.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
