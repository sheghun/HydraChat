package hydrachat

import (
	"Hydra/hlogger"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var logger = hlogger.GetInstance()

func Run(connection string) error {
	l, err := net.Listen("tcp", connection)
	r := CreateRoom("Hydrachat")

	if err != nil {
		logger.Println("Error connecting to chat client", err)
		return err
	}

	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch

		l.Close()
		fmt.Println("Closing tcp connection")
		close(r.Quit)

		if r.ClientsCount() > 0 {
			<-r.Msgch
		}
		os.Exit(0)

	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Println("Error accepting connection from chat client", err)
		}
	go handleConnection(r, conn)
	}

	return err

}

func handleConnection(r *room, conn net.Conn) {
	logger.Println("Received request from client", conn.RemoteAddr())
	r.AddClient(conn)
}
