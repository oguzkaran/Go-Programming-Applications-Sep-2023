package app

import (
	"Sender/csd/console"
	"fmt"
	"net"
	"os"
)

func getUDPAddress(address string) *net.UDPAddr {
	udp, e := net.ResolveUDPAddr("udp", address)

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "ResolveUDPAddr:%s\n", e.Error())
		os.Exit(1)
	}

	return udp
}

func getUDPConnection(udp *net.UDPAddr) *net.UDPConn {
	con, e := net.DialUDP("udp", nil, udp)
	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "DialUDP:%s\n", e.Error())
		os.Exit(1)
	}

	return con
}

func send(con *net.UDPConn) {
	for {
		s := console.ReadString("Input string:")

		if s == "quit" {
			break
		}

		data := []byte(s)

		_, _ = con.Write(data)
	}
}

func Run() {
	console.CheckLengthEquals(3, len(os.Args), "usage: ./sender <host> <port number>\n")
	address := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])
	udp := getUDPAddress(address)
	con := getUDPConnection(udp)
	send(con)
}
