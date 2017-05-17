package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Server struct{}

func (server *Server) Negate(input int64, reply *int64) error {
	*reply = -1 * input
	return nil
}

func server() {
	rpc.Register(new(Server)) // register the a "Server" object for remote-procdedure-call
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn) // serving the connection just established.
	}
}

func client() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:9999") // Try to make a connection on localhost and port 9999
	if err != nil {                                // connection failed.
		fmt.Println(err)
		return
	}
	// Try to negate an integer
	var result int64
	err = conn.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", result)
	}

}

func main() {
	go server()
	go client()
	var input string
	fmt.Scanln(&input)
}
