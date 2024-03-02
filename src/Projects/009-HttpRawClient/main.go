package main

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

func main() {
	const headStr = "HEAD /HTTP/1.0\r\n\r\n"
	checkLengthEquals(2, len(os.Args), "wrong number of arguments!...")
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)

	checkError(err)
	fmt.Println(tcpAddr)

	con, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = con.Write([]byte(headStr))
	checkError(err)
	result, err := io.ReadAll(con)

	fmt.Println(string(result))

}
