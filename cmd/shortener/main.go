package main

import (
	"fmt"
	"net/http"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/handler"

	_ "github.com/lib/pq"
)

func main() {
	config := app.NewConfig()
	mux := handler.InitHandlers()

	srv := http.Server{
		Addr:    config.ServerAddress,
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Print(err)
	}
}
