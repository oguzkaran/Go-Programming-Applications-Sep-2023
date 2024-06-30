package app

import (
	"Server/app/shared"
	"Server/csd/console"
	"errors"
	"fmt"
	"math/rand/v2"
	"net"
	"net/rpc"
	"os"
)

type RandomNumberGenerator int

func (rg *RandomNumberGenerator) GenerateNumber(info *shared.RandomNumberInfo, result *int) error {
	if info.Min >= info.Bound {
		return errors.New(fmt.Sprintf("%d must be less than %d\n", info.Min, info.Bound))
	}

	*result = rand.IntN(info.Bound-info.Min) + info.Min
	return nil
}

func registerGenerator(gen *RandomNumberGenerator) {
	if e := rpc.Register(gen); e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "registerGenerator: %s\n", e.Error())
		os.Exit(1)
	}
}

func getTCPAddr() *net.TCPAddr {
	tcp, e := net.ResolveTCPAddr("tcp", ":"+os.Args[1])

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "getTCPAddr: %s\n", e.Error())
		os.Exit(1)
	}

	return tcp
}

func getTCPListener(addr *net.TCPAddr) *net.TCPListener {
	listener, e := net.ListenTCP("tcp", addr)
	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "getTCPListener: %s\n", e.Error())
		os.Exit(1)
	}

	return listener
}

func Run() {
	console.CheckLengthEquals(2, len(os.Args), "usage: ./server <port number>")
	gen := new(RandomNumberGenerator)
	registerGenerator(gen)
	tcp := getTCPAddr()
	listener := getTCPListener(tcp)

	defer func(tcpListener *net.TCPListener) {
		_ = listener.Close()
	}(listener)

	for {
		con, e := listener.Accept()
		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Accept: %s\n", e.Error())
			continue
		}

		rpc.ServeConn(con)
	}
}
