package models

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	utils "loadbalancer/utils"
)

type Server struct {
	Addr  string
	Proxy *httputil.ReverseProxy
}

func (s *Server) Address() string {
	return s.Addr
}

func (s *Server) IsAlive() bool {
	return true
}

func (s *Server) Serve(rw http.ResponseWriter, req *http.Request) {
	s.Proxy.ServeHTTP(rw, req)
}

func NewServer(addr string) Server {
	serverUrl, err := url.Parse(addr)

	fmt.Printf("Parsed url: %s\n", serverUrl)
	utils.HandleErr(err)

	return Server{
		Addr:  addr,
		Proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}
