package cristian

import (
	"log"
	"net/rpc"
	"time"
)

func Sync(serverAddr string) (time.Time, error) {
	client, err := rpc.Dial("tcp", serverAddr)
	CheckError(err)
	T1 := time.Now()
	reply := CristianReply{T1,T1}
	err = client.Call("CristianServer.Respond", struct{}{}, &reply)
	if err != nil {
		return time.Now(), err
	}
	T4 := time.Now()
	return reply.T3.Add((T4.Sub(T1)-(reply.T3.Sub(reply.T2)))/2), nil
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
