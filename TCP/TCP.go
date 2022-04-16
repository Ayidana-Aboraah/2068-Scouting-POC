package TCP

import (
	"bufio"
	"log"
	"net"
	"strings"
)

var conn net.Conn

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
			<-done
			return
		}
	}()

	<-done
}

func ShutDown() {
	connection, _ := net.Dial("tcp", ":9500")
	connection.Write([]byte("shutdown\n"))
	connection.Close()
}

func FindIP() string {
	conn, error := net.Dial("udp", "8.8.8.8:80")
	if error != nil {
		log.Println(error)
	}

	ipAddress := conn.LocalAddr().(*net.UDPAddr)
	conn.Close()

	i := strings.Index(ipAddress.String(), ":")
	return ipAddress.String()[:i] // + ":9500"
}

func ConnectToTCP(ip string) error {
	connection, err := net.Dial("tcp", ip+":9500")
	conn = connection

	connection.Write([]byte("Comp list\n"))
	message, _ := bufio.NewReader(connection).ReadString('\n')

	compKeys = strings.Split(message, "¶")
	return err
}

func DisconnectTCP() {
	if conn != nil {
		conn.Close()
	}
}
