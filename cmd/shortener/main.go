package main

import (
	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/handler"

	_ "github.com/lib/pq"
)

func main() {
	config := app.NewConfig()
	handler.NewHandler(config)
}
