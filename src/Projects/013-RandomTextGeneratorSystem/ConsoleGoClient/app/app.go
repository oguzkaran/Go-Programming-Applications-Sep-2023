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
	_, _ = fmt.Fprintf(os.Stderr, "%s\n", message)
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

	var buf [1024]byte

	n, err := socket.Read(buf[0:])
	checkError(err)
	fmt.Println(string(buf[0:n]))

	_, err = socket.Write([]byte(fmt.Sprintf("My Time:%s", time.Now().String())))
	checkError(err)
	checkError(socket.Close())
}
