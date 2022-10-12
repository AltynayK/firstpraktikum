package main

import (
	"context"
	"fmt"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/handler"
	_ "github.com/lib/pq"
)

func main() {
	config := app.NewConfig()
	s := handler.NewHandler(config)
	s.Run(config)
	_, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		fmt.Println("Run cancel in defer...")
	}()

}
