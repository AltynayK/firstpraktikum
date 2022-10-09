package main

import (
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
	tooLate := make(chan struct{})
	proCh := make(chan string)

	go func() {
		for {
			fmt.Println("working")
			time.Sleep(1 * time.Second)
			select {
			case <-tooLate:
				fmt.Println("stopped")
				return
			case proCh <- "processed":
			default:
			}
			fmt.Println("done here")

		}
	}()
	select {
	case proc := <-proCh:
		fmt.Println(proc)
	case <-time.After(1 * time.Second):
		fmt.Println("too late")
		close(tooLate)
	}

	time.Sleep(4 * time.Second)
	fmt.Print("finish\n")
}
