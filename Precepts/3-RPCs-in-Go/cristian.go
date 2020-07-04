package main

import (
	"fmt"
	"time"
	"cristian"
)

func main() {
	serverAddr := "localhost:8888"
	server := cristian.CristianServer{serverAddr}
	server.Listen()
	synced_t, err := cristian.Sync(serverAddr)
	cristian.CheckError(err)
	fmt.Printf("%d nanoseconds slower\n", synced_t.Sub(time.Now()).Nanoseconds())
}