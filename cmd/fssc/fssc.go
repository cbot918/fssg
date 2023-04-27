package main

import (
	"github.com/cbot918/fssg/fssc"
)

const (
	url = "localhost:8888"
)

func main() {
	c := fssc.NewFssc(url)
	c.Run()
}
