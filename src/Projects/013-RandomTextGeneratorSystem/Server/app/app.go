/*
Client -> Server
count -> uint64
origin -> uint32
bound -> uin32

Server -> Client
success -> 0 ya da 1

if success
for count

	send random text
*/
package app

import (
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

func handleClient(socket net.Conn) {
	fmt.Printf("Client connected:%v\n", socket.RemoteAddr())

	countBuf := make([]byte, 8)
	originBuf := make([]byte, 4)
	boundBuf := make([]byte, 4)

	n, err := socket.Read(countBuf)

	if err != nil || n != len(countBuf) {
		_, _ = socket.Write([]byte{0})
		return
	}
	n, err = socket.Read(originBuf)

	if err != nil || n != len(originBuf) {
		_, _ = socket.Write([]byte{0})
		return
	}

	n, err = socket.Read(boundBuf)

	if err != nil || n != len(boundBuf) {
		_, _ = socket.Write([]byte{0})
		return
	}
	var count uint64
	var origin uint32
	var bound uint32

	count = binary.NativeEndian.Uint64(countBuf)
	origin = binary.NativeEndian.Uint32(originBuf)
	bound = binary.NativeEndian.Uint32(boundBuf)

	fmt.Printf("Count:%d, Origin:%d, Bound:%d\n", count, origin, bound)

	_, _ = socket.Write([]byte{1})

	_ = socket.Close()
}

func Run() {
	checkLengthEquals(2, len(os.Args), "usage: ./server <port number>")
	service := fmt.Sprintf(":%s", os.Args[1])

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	serverSocket, err := net.ListenTCP("tcp", tcpAddr)

	fmt.Printf("RandomTextGeneratorServer is waiting for a client on port :%s\n", os.Args[1])
	for {
		socket, err := serverSocket.Accept()

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			continue
		}

		go handleClient(socket)
	}
}
