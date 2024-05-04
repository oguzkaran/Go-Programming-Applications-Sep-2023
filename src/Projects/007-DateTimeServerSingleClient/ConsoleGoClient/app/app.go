package app

import (
	"fmt"
	"io"
	"net"
	"os"
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

func Run() {
	checkLengthEquals(3, len(os.Args), "usage: ./client <host> <port number>")

	service := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	socket, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	result, err := io.ReadAll(socket)
	checkError(err)
	fmt.Printf(string(result))
	socket.Close()
}
