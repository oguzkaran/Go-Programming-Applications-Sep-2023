package tcp

import (
	"fmt"
	"net"
	"os"
)

func StartServer(address, initPrompt, clientErrorPrompt string, clientCallback func(socket net.Conn)) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)

	if err != nil {
		panic(err)
		return
	}

	serverSocket, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		panic(err)
		return
	}

	fmt.Print(initPrompt)

	for {
		socket, err := serverSocket.Accept()

		if err != nil {
			if clientErrorPrompt != "" {
				_, _ = fmt.Fprintf(os.Stderr, "%s:%s", clientErrorPrompt, err.Error())
			}
			continue
		}

		go clientCallback(socket)
	}
}
