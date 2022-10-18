package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/handler"
	_ "github.com/lib/pq"
)

func main() {
	config := app.NewConfig()
	s := handler.NewHandler(config)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := s.Run(ctx, config); err != nil {
		log.Fatal(err)
	}

}
