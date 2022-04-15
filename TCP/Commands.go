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
		case "SendT":
			log.Println(FromBytes(s.Bytes()[6:]))
			conn.Write(append(s.Bytes()[6:], []byte("\n")...))
		// case "GET":
		// 	get(str[1:], conn)
		// case "SET":
		// 	set(str[1:], conn)
		default:
			conn.Write(InvalidCommand)
		}

		conn.Write([]byte("\n>"))
		// }()
	}
}

func competition(cmd []string, conn net.Conn) {
	switch cmd[0] {
	case "list":
		for _, comp := range competitions {
			conn.Write([]byte(comp))
		}
		conn.Write([]byte("\n"))
	case "find":
		for _, comp := range competitions {
			if cmd[1] != comp {
				continue
			}

			//Send back the form for that competitons
			conn.Write([]byte(comp + "\n"))
		}
	case "add":
		for _, comp := range competitions {
			if cmd[1] == comp {
				conn.Write([]byte("!"))
				break
			}
		}
		competitions = append(competitions, cmd[1])
	}
}

// func set(cmd []string, conn net.Conn) {

// 	if len(cmd) < 2 {
// 		conn.Write(InvalidCommand)
// 		return
// 	}

// 	key := cmd[0]
// 	val := cmd[1]

// 	c.Lock()
// 	c.data[key] = val
// 	c.Unlock()

// 	conn.Write([]byte("Added"))
// }

// func get(cmd []string, conn net.Conn) {

// 	if len(cmd) < 1 {
// 		conn.Write(InvalidCommand)
// 		return
// 	}

// 	val := cmd[0]

// 	c.RLock()
// 	ret, ok := c.data[val]
// 	c.RUnlock()

// 	if !ok {
// 		conn.Write([]byte("Nil"))
// 		return
// 	}

// 	conn.Write([]byte(ret))
// }
