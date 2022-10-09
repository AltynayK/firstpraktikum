package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/handler"
	_ "github.com/lib/pq"
)

func main() {
	config := app.NewConfig()
	s := handler.NewHandler(config)
	s.Run(config)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("missed signal")
	case <-ctx.Done():
		stop()
		fmt.Println("signal received")
	}

}
