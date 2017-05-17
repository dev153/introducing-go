package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go HandleServerConnection(conn)
	}
}

func HandleServerConnection(conn net.Conn) {
	var msg string
	err := gob.NewDecoder(conn).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	conn.Close()
}

//=============================================================================

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(conn).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	conn.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
