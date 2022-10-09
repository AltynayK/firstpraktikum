package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/AltynayK/firstpraktikum/internal/app"
	"github.com/AltynayK/firstpraktikum/internal/handler"
	_ "github.com/lib/pq"
)

func main() {
	config := app.NewConfig()
	s := handler.NewHandler(config)
	s.Run(config)
	parent, pCancel := context.WithCancel(context.Background())
	child, _ := context.WithCancel(parent)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go work(wg, child)
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	defer signal.Stop(c)
	select {
	case <-c:
		pCancel()
		fmt.Println("Waiting everyone to finish...")
		wg.Wait()
		fmt.Println("Exiting")
		os.Exit(0)
	}

}
func work(wg *sync.WaitGroup, ctx context.Context) {
	done := false
	wg.Add(1)
	for !done {
		fmt.Println("Doing something...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			done = true
		default:
		}
	}
	wg.Done()
}
