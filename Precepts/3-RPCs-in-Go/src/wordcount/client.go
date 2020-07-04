package wordcount

import (
	"log"
	"net/rpc"
)

func MakeRequest(input string, serverAddr string) (map[string]int, error) {
	client, err := rpc.Dial("tcp", serverAddr)
	CheckError(err)
	args := WordCountRequest{input}
	reply := WordCountReply{make(map[string]int)}
	err = client.Call("WordCountServer.Compute", args, &reply)
	if err != nil {
		return nil, err
	}
	return reply.Counts, nil
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

   