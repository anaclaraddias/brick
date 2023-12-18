package main

import "github.com/anaclaraddias/brick/adapter/http"

func main() {
	httpserver := http.NewServer()
	if err := httpserver.Start(); err != nil {
		panic(err)
	}
}
