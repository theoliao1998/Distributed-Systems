package cristian

import (
	"time"
	"net"
	"net/rpc"
)

type CristianServer struct {
	Addr string
}

type CristianReply struct {
	T2 time.Time
	T3 time.Time
}

func (*CristianServer) Respond(_ struct{}, reply *CristianReply) error {
	reply.T2 = time.Now()
	time.Sleep(100 * time.Millisecond)
	reply.T3 = time.Now()
	return nil
}

func (server *CristianServer) Listen() {
	rpc.Register(server)
	listener, err := net.Listen("tcp", server.Addr)
	CheckError(err)
	go func() {
		for {
			rpc.Accept(listener)
		}
	}()
}