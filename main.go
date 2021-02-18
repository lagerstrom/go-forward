package main

import (
	"github.com/lagerstrom/go-forward/forwarder"
)

func main() {
	forwarder.ParseConfig()
	//done := make(chan bool)
	go forwarder.StartServer()
	//http://stackoverflow.com/questions/9543835/how-best-do-i-keep-a-long-running-go-program-running
	select {} // block forever, this program exits by ctrl+c :)
}
