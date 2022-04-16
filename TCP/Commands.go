package TCP

import (
	"bufio"
	"log"
	"net"
	"strings"
)

var InvalidCommand = []byte("Invalid Command")

func HandleConnection(conn net.Conn, done chan bool) {
	defer conn.Close()

	s := bufio.NewScanner(conn)

	for s.Scan() {

		data := s.Text()

		if data == "exit" {
			return
		}

		if data == "shutdown" {
			log.Println("Shutting Down")
			done <- true
			return
		}

		// handleCommand(data, conn)
		// func() {
		str := strings.Split(data, " ")

		if len(str) <= 0 {
			conn.Write(InvalidCommand)
			return
		}

		switch str[0] { //Checking the Command
		case "Test":
			conn.Write([]byte("Dekimakura\n"))
			return
		case "Comp":
			competition(str[1:], conn)
		case "Submit":
			database.competitions[str[1]] = []Form{} //For Testing

			submission := FromBytes(s.Bytes()[len(str[1])+8:])
			database.competitions[str[1]] = append(database.competitions[str[1]], submission)
			conn.Write(ToBytes(database.competitions[str[1]][0]))
			// conn.Write(s.Bytes()[8+len(str[1]):])
		default:
			conn.Write([]byte("!")) //This will act as our basic error message
		}

		conn.Write([]byte("\n>"))
		// }()
	}
}

func competition(cmd []string, conn net.Conn) {
	switch cmd[0] {
	case "list":
		var temp string
		for comp := range compTemplates {
			temp += comp + "¶"
		}
		conn.Write([]byte(temp + "\n"))
	default:
		if form, found := compTemplates[cmd[1]]; found {
			conn.Write(ToBytes(form))
			return
		}

		conn.Write([]byte("!f"))
	}
}
