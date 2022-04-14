package TCP

import (
	"log"
	"net"
)

func StartTCP() {
	listener, err := net.Listen("tcp", ":9500")
	if err != nil {
		// t.Error(err)
		log.Println(err)
	}

	done := make(chan bool)

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
			go HandleConnection(conn, done)
		}
	}()

	<-done

	listener.Close()
}
