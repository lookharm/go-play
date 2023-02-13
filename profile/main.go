package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	transport := &http.Transport{
		DisableKeepAlives: true,
	}
	var (
		address = "http://localhost:3001"
		id      = "GBvfOME3DZ"
		status  = "up"
		msg     = ""
	)

	for {
		url := fmt.Sprintf("%s/api/push/%s?status=%s&msg=%s", address, id, status, msg)
		c := http.Client{Transport: transport}
		res, _ := c.Get(url)
		res.Body.Close()
		time.Sleep(10 * time.Millisecond)
	}
}
