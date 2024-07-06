package app

import (
	"Server/app/shared/impl"
	"Server/csd/console"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

func registerGenerator(gen *impl.RandomNumberGenerator) {
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
	console.CheckLengthEquals(2, len(os.Args), "usage: ./rpc_server <port number>")
	gen := new(impl.RandomNumberGenerator)
	registerGenerator(gen)
	tcp := getTCPAddr()
	listener := getTCPListener(tcp)

	defer func(tcpListener *net.TCPListener) {
		_ = listener.Close()
	}(listener)

	_ = console.WriteLine("RPC server is waiting for a client on :%s", os.Args[1])

	for {
		con, e := listener.Accept()
		if e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Accept: %s\n", e.Error())
			continue
		}

		go rpc.ServeConn(con) //Bu işlemde belli bir süre bağlı kalıp işlem yapmayan client'ların bağlantıları kesilmeli
	}
}
