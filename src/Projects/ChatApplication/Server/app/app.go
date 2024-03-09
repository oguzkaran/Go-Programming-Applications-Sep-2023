package app

import (
	"Server/csd/net/tcp"
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

func registerServerCallback(socket net.Conn) {
	defer socket.Close()

}

func loginServerCallback(socket net.Conn) {
	defer socket.Close()
}

func Run() {
	checkLengthEquals(3, len(os.Args), "usage: ./server <register port> <login port>")
	registerAddress := fmt.Sprintf(":%s", os.Args[1])
	loginAddress := fmt.Sprintf(":%s", os.Args[2])
	go tcp.StartServer(registerAddress, "RegisterServer\n", "", registerServerCallback)
	tcp.StartServer(loginAddress, "LoginServer\n", "", loginServerCallback)
}
