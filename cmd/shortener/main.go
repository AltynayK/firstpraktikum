package main

import (
	"context"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/handler"
	_ "github.com/lib/pq"
)

func main() {
	config := app.NewConfig()
	s := handler.NewHandler(config)
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	s.Run(config)

}
