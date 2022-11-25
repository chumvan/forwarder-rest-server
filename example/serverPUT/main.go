package main

import (
	"fmt"
	"sync"

	models "github.com/chumvan/confdb/models"
	router "github.com/chumvan/forwarder-rest-server/routers"
)

func main() {
	usersChan := make(chan []models.User, 1)
	r := router.SetupRouter(usersChan)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		r.Run()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for users := range usersChan {
			fmt.Printf("users are %v\n", users)
		}
	}()

	wg.Wait()
}
