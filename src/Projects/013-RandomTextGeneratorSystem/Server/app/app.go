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
	"Server/csd/str"
	"encoding/binary"
	"fmt"
	"math/rand"
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

func send(conn net.Conn, buf []byte) (int, error) {
	count := len(buf)
	for {
		n, err := conn.Write(buf)
		if err != nil {
			return 0, err
		}
		count -= n
		if count == 0 {
			break
		}
	}

	return len(buf), nil
}

func receive(conn net.Conn, buf []byte) (int, error) {
	size := len(buf)
	count := 0
	data := make([]byte, size)

	for {
		n, err := conn.Read(data)
		if err != nil {
			return 0, err
		}

		for i := 0; i < len(data); i++ {
			buf[i+count] = data[i]
		}
		count += n

		if count == size {
			break
		}
	}

	return count, nil
}

func handleClient(socket net.Conn) {
	defer socket.Close()
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

	_, err = socket.Write([]byte{1})
	if err != nil {
		return
	}

	dataLen := make([]byte, 4)
	for i := uint64(0); i < count; i++ {
		text := str.GenerateRandomTextEN(rand.Intn(int(bound)-int(origin)) + int(origin))

		binary.NativeEndian.PutUint32(dataLen, uint32(len(text)))
		n, err := socket.Write(dataLen)
		n, err = socket.Write([]byte(text))
		fmt.Printf("length of %s is %d, number of written data:%d\n", text, len(text), n)
		if err != nil {
			return
		}
	}
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
