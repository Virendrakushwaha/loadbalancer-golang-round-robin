package models

import (
	"fmt"
	"net/http"
)

type LoadBalancer struct {
	Port            string
	RoundRobinCount int
	Servers         []Server
}

func (lb *LoadBalancer) GetNextAvailableServer() Server {
	server := lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	fmt.Printf("current count: %v\n", lb.RoundRobinCount)
	for !server.IsAlive() {
		lb.RoundRobinCount++
		server = lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	}
	lb.RoundRobinCount++
	return server
}

func (lb *LoadBalancer) ServeProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.GetNextAvailableServer()
	fmt.Printf("Forwarding request to address : %q\n", targetServer.Address())
	targetServer.Serve(rw, req)
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		RoundRobinCount: 0,
		Servers:         servers,
	}
}
