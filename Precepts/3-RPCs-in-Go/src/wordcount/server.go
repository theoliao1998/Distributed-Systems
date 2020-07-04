package wordcount

import (
	"strings"
	"net"
	"net/rpc"
)

type WordCountServer struct {
	Addr string
}

type WordCountRequest struct {
	Input string
}

type WordCountReply struct {
	Counts map[string]int
}

func (*WordCountServer) Compute(request WordCountRequest, reply *WordCountReply) error {
	counts := make(map[string]int)
	input := request.Input
	tokens := strings.Fields(input)
	for _, t := range tokens {
		counts[t] += 1
	}
	reply.Counts = counts
	return nil
}

func (server *WordCountServer) Listen() {
	rpc.Register(server)
	listener, err := net.Listen("tcp", server.Addr)
	CheckError(err)
	go func() {
		for {
			rpc.Accept(listener)
		}
	}()
}