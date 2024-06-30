package app

import (
	"Receiver/csd/console"
	"fmt"
	"net"
	"os"
)

const network = "udp"
const dgramSize = 2048

func getUDPAddress(address string) *net.UDPAddr {
	udp, e := net.ResolveUDPAddr(network, address)

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "getUDPAddress:%s\n", e.Error())
		os.Exit(1)
	}

	return udp
}

func getUDPConnection(udp *net.UDPAddr) *net.UDPConn {
	con, e := net.ListenUDP(network, udp)

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "getUDPConnection:%s\n", e.Error())
		os.Exit(1)
	}

	return con
}

func receive(con *net.UDPConn) {
	buf := make([]byte, dgramSize)

	for {
		fmt.Printf("Receiver is waiting for a client on port:%s\n", os.Args[1])
		result, udpAddr, e := con.ReadFromUDP(buf)

		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "ReadFromUDP:%s\n", e.Error())
			os.Exit(1)
		}
		handleClient(con, udpAddr, buf, result)
	}
}

func handleClient(con *net.UDPConn, addr *net.UDPAddr, buf []byte, n int) {
	data := buf[0:n]
	str := string(data)

	fmt.Printf("'%s' received from %s:%s\n", str, addr.IP.String(), addr.AddrPort().String())
}

func Run() {
	console.CheckLengthEquals(2, len(os.Args), "usage: ./receiver <port number>\n")
	address := ":" + os.Args[1]
	udp := getUDPAddress(address)
	con := getUDPConnection(udp)

	defer func(con *net.UDPConn) {
		_ = con.Close()
	}(con)
	receive(con)
}
