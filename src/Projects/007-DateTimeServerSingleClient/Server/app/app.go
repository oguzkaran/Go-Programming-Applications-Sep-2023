package app

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkLengthEquals(len, argsLen int, message string) {
	if argsLen != len {
		exitFailure(message)
	}
}

func exitFailure(message string) {
	fmt.Fprintf(os.Stderr, "%s\n", message)
	os.Exit(1)
}

func checkError(err error) {
	if err != nil {
		exitFailure(err.Error())
	}
}

func handleClient(socket net.Conn) {
	fmt.Printf("Client connected:%v\n", socket.RemoteAddr())
	nowStr := time.Now().String()
	socket.Write([]byte(nowStr))
	checkError(socket.Close())
}

func Run() {
	checkLengthEquals(2, len(os.Args), "usage: ./server <port number>")
	service := fmt.Sprintf(":%s", os.Args[1])

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	serverSocket, err := net.ListenTCP("tcp", tcpAddr)

	for {
		fmt.Printf("Date Time Server is waiting for a client on port :%s\n", os.Args[1])
		socket, err := serverSocket.Accept()

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			continue
		}

		handleClient(socket)
	}
}
