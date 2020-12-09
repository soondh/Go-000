//simple server struct without handler
package main

import "net/http"

type server struct {
	*http.Server
	name string
}

func newServer(name, addr string) server {
	return server{
		&http.Server{
			Addr:    addr,
		},
		name,
	}
}
