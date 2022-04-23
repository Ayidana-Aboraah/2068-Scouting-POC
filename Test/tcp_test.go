package test

import (
	"2068_Scouting/TCP"
	"bufio"
	"log"
	"net"
	"strings"
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
	timer := time.NewTimer(5 * time.Second)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("Accept Error", err)
				continue
			}

			log.Println("Accepted ", conn.RemoteAddr())

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

func TestSendTCP(t *testing.T) {
	go TCP.StartTCP()

	connection, err := net.Dial("tcp", ":9500")
	if err != nil {
		t.Error(err)
	}
	defer connection.Close()

	connection.Write([]byte("shutdown\n"))
}

//Make a recieve function

func TestRecieveTCP(t *testing.T) {
	time.Sleep(1 * time.Second)
	go TCP.StartTCP()

	connection, err := net.Dial("tcp", ":9500")
	if err != nil {
		t.Error(err)
	}
	defer connection.Close()

	connection.Write([]byte("Test\n"))

	message, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		t.Error(err)
	}

	t.Log("Message from server: " + message)

	if !strings.Contains(message, "Dekimakura") {
		t.Error("Results not matching")
	}
}

func TestSubmit(t *testing.T) {
	time.Sleep(1 * time.Second)
	go TCP.StartTCP()

	connection, err := net.Dial("tcp", ":9500")
	if err != nil {
		t.Error(err)
	}
	defer connection.Close()

	form := TCP.Form{
		Team:      2022,
		Questions: []string{"Location Of My Mom?"},
		Answers:   []string{"In My Bed"},
	}

	connection.Write(append([]byte("Submit Test "), TCP.ToBytes(form)...))

	scanner := bufio.NewScanner(connection)
	if err != nil {
		t.Error(err)
	}

	scanner.Scan()

	t.Log(scanner.Text())

	newForm := TCP.FromBytes(scanner.Bytes(), false)

	CompareForm(form, newForm, t)

	t.Log(form)
	t.Log(newForm)
}
