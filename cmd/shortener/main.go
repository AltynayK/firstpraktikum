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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := s.Run(ctx, config); err != nil {
		fmt.Print(err)
	}

}
