package app

import (
	"Server/console"
	"encoding/binary"
	"fmt"
	"net"
	"os"
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

	countBuf := make([]byte, 8)
	originBuf := make([]byte, 4)
	boundBuf := make([]byte, 4)

	count := uint64(console.ReadInt("Input count:", ""))
	origin := uint32(console.ReadInt("Input origin:", ""))
	bound := uint32(console.ReadInt("Input bound:", ""))

	binary.NativeEndian.PutUint64(countBuf, count)
	binary.NativeEndian.PutUint32(originBuf, origin)
	binary.NativeEndian.PutUint32(boundBuf, bound)

	_, err = socket.Write(countBuf)
	checkError(err)
	_, err = socket.Write(originBuf)
	checkError(err)
	_, err = socket.Write(boundBuf)
	checkError(err)
	status := make([]byte, 1)
	_, err = socket.Read(status)
	checkError(err)

	if status[0] == 1 {
		fmt.Println("Success")
	} else {
		fmt.Println("Fail")
	}

	checkError(socket.Close())
}
