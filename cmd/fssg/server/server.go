package main

import (
	"github.com/cbot918/fssg/src/server"
)

const (
	url = "localhost:8888"
)

func main() {
	fss := server.NewFssg(url)
	fss.Run()
}
