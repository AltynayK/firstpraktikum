package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/handler"
	_ "github.com/lib/pq"
)

func main() {
	config := app.NewConfig()
	s := handler.NewHandler(config)

	ctx, cancel := context.WithCancel(context.Background())
	go handleSignals(cancel)
	if err := s.Run(ctx, config); err != nil {
		fmt.Print(err)
	}
}
func handleSignals(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	for {
		sig := <-sigCh
		switch sig {
		case os.Interrupt:
			cancel()
			return
		}
	}
}
