package TCP

import (
	"log"
	"net"
)

func StartTCP() {
	listener, err := net.Listen("tcp", ":9500")
	if err != nil {
		log.Println(err)
	}
	defer listener.Close()

	done := make(chan bool)
	defer close(done)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("Accept Error", err)
				continue
			}

			log.Println("Accepted ", conn.RemoteAddr())

			//create a routine dont block
			go HandleConnection(conn, done)
		}
	}()

	<-done
}
