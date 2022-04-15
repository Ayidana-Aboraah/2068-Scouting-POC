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
	timer := time.NewTimer(10 * time.Second)
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

func TestRecieveTCP(t *testing.T) {
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

func TestSendAndRecieve(t *testing.T) {
	go TCP.StartTCP()

	connection, err := net.Dial("tcp", ":9500")
	if err != nil {
		t.Error(err)
	}
	defer connection.Close()

	form := TCP.Form{
		Team:      2020,
		Questions: []string{"How old are you?"},
		Answers:   []string{"Your Mom"},
	}

	connection.Write(append([]byte("SendT "), append(TCP.ToBytes(form), []byte("\n")...)...))

	scaner := bufio.NewScanner(connection)
	if err != nil {
		t.Error(err)
	}

	scaner.Scan()

	newForm := TCP.FromBytes(scaner.Bytes())

	//Compare the 2
	if form.Team != newForm.Team {
		t.Error("Teams don't match, data corruption active.")
	}

	for i := range form.Questions {
		if form.Questions[i] != newForm.Questions[i] {
			t.Error("The questions don't match up")
			t.Errorf("form: %v, new form: %v", form.Questions[i], newForm.Questions[i])
		}

		if form.Answers[i] != newForm.Answers[i] {
			t.Error("The answers don't match up")
			t.Errorf("form: %v, new form: %v", form.Answers[i], newForm.Answers[i])
		}
	}

	t.Log(form)
	t.Log(newForm)
}
