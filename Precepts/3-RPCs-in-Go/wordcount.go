package main

import (
	"fmt"
	"wordcount"
)

func main() {
	serverAddr := "localhost:8888"
	server := wordcount.WordCountServer{serverAddr}
	server.Listen()
	input1 := "hello I am good hello bye bye bye bye good night hello"
	wc, err := wordcount.MakeRequest(input1, serverAddr)
	wordcount.CheckError(err)
	fmt.Printf("Result: %v\n", wc)
}