package hydrachat

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	go func() {
		t.Log("Starting Hydra chat server... ")
		if err := Run(":2000"); err != nil {
			t.Error("could not start chat server", err)
		} else {
			t.Log("Started Hydra chat server...")
		}
	}()

	time.Sleep(1 * time.Second)
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))

	t.Logf("Hello %s, connnecting to the hydra chat system... \n", name)
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		t.Fatal("could not connect to hydra chat system", err)
	}
	t.Log("Connected to hydra chat system")
	name += ":"
	defer conn.Close()

	msgCh := make(chan string)

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			reMsg := scanner.Text()
			sentMsg := <-msgCh
			if strings.Compare(reMsg, sentMsg) != 0 {
				t.Errorf("Chat message %s does not match %s", reMsg, sentMsg)
			}
		}
	}()

	for i := 0; i <= 10; i++ {
		msgBody := fmt.Sprintf("RandomMessage %d", rand.Intn(400))
		msg := name + msgBody
		_, err = fmt.Fprintf(conn, msg+"\n")
		if err != nil {
			t.Error(err)
			return
		}
		msgCh <- msg
	}

}
