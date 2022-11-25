package main

import (
	"fmt"
	"net"
	"net/url"
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
			fmt.Println("received users")
			for _, u := range users {
				fmt.Printf("full string: %v\n", u)
				url, err := url.Parse(u.EntityUrl.String())
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("user: %s\n", url.User.Username())
				host, port, err := net.SplitHostPort(url.Host)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Printf("host: %s\n", host)
				fmt.Printf("port: %s\n", port)
			}
		}
	}()

	wg.Wait()
}
