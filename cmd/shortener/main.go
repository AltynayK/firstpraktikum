package main

import (
	"context"
	"fmt"
	"time"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/handler"
	_ "github.com/lib/pq"
)

func main() {
	config := app.NewConfig()
	s := handler.NewHandler(config)

	s.Run(config)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	doWork(ctx)
}

func doWork(ctx context.Context) {
	newCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	fmt.Print("starting working ...")

	for {
		select {
		case <-newCtx.Done():
			fmt.Print("ctx done")
			return
		default:
			fmt.Print("working...")
			time.Sleep(1 * time.Second)
		}
	}
}
