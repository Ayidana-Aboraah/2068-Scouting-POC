package test

import (
	"2068_Scouting/TCP"
	"log"
	"net"
	"testing"
)

func TestTCP(t *testing.T) {
	listener, err := net.Listen("tcp", ":9500")
	if err != nil {
		t.Error(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept Error", err)
			continue
		}

		log.Println("Accepted ", conn.RemoteAddr())
		conn.Write([]byte(">"))

		//create a routine dont block
		go TCP.HandleConnection(conn)
	}
}
