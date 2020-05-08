package rpcSupper

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServerRpc(host string, service interface{}) error {
	err := rpc.Register(service)
	if err != nil {
		return err
	}
	listener, err := net.Listen("tcp", host)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	client := jsonrpc.NewClient(conn)
	return client, nil
}
