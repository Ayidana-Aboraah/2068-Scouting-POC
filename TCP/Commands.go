package TCP

import (
	"bufio"
	"net"
	"strings"
	"sync"
)

type cache struct {
	data map[string]string
	*sync.RWMutex
}

var c = cache{data: make(map[string]string), RWMutex: &sync.RWMutex{}}
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
			conn.Write([]byte("Shutting Down..."))
			done <- true
			return
		}

		handleCommand(data, conn)
	}
}

func handleCommand(inp string, conn net.Conn) {

	str := strings.Split(inp, " ")

	if len(str) <= 0 {
		conn.Write(InvalidCommand)
		return
	}

	switch str[0] { //Checking the Command
	case "GET":
		get(str[1:], conn)
	case "SET":
		set(str[1:], conn)
	default:
		conn.Write(InvalidCommand)
	}

	conn.Write([]byte("\n>"))
}

func set(cmd []string, conn net.Conn) {

	if len(cmd) < 2 {
		conn.Write(InvalidCommand)
		return
	}

	key := cmd[0]
	val := cmd[1]

	c.Lock()
	c.data[key] = val
	c.Unlock()

	conn.Write([]byte("OK"))
}

func get(cmd []string, conn net.Conn) {

	if len(cmd) < 1 {
		conn.Write(InvalidCommand)
		return
	}

	val := cmd[0]

	c.RLock()
	ret, ok := c.data[val]
	c.RUnlock()

	if !ok {
		conn.Write([]byte("Nil"))
		return
	}

	conn.Write([]byte(ret))
}