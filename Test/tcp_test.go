package test

import (
	"2068_Scouting/TCP"
	"log"
	"net"
	"testing"
	"time"
)

func TestStartTCP(t *testing.T) {
	listener, err := net.Listen("tcp", ":9500")
	if err != nil {
		t.Error(err)
	}
	defer listener.Close()

	done := make(chan bool)
	timer := time.NewTimer(10 * time.Second)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("Accept Error", err)
				continue
			}

			log.Println("Accepted ", conn.RemoteAddr())
			conn.Write([]byte(">"))

			//create a routine dont block
			go TCP.HandleConnection(conn, done)
		}
	}()

	select {
	case <-done:
		return
	case <-timer.C:
		return
	}
}
