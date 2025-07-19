package main

import (
	"fmt"
	models "loadbalancer/models"
	"net/http"
)

type IServer interface {
	Address() string
	IsAlive() bool
	Server(rw http.ResponseWriter, req *http.Request)
}

func main() {

	// create list of servers
	servers := []models.Server{
		models.NewServer("https://github.com"),
		models.NewServer("https://google.com"),
		models.NewServer("https://linkedin.com"),
	}

	// create load balancer
	lb := models.NewLoadBalancer(":8080", servers)

	// redirect request
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	}

	// http entry point
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Serving requests at localhost: %v", 8080)

	// Server http
	http.ListenAndServe(lb.Port, nil)
}
