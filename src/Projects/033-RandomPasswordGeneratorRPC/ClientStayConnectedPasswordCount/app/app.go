package app

import (
	"Client/csd/console"
	"net/rpc"
	"os"
)

func connect() *rpc.Client {
	c, e := rpc.Dial("tcp", os.Args[1]+":"+os.Args[2])

	if e != nil {
		_, _ = os.Stderr.WriteString("invalid address")
		os.Exit(1)
	}

	return c
}

func callGenerate(client *rpc.Client, count *int) {
	var result string

	e := client.Call("RandomPasswordGenerator.GeneratePassword", count, &result)
	if e != nil {
		_, _ = os.Stderr.WriteString("RPC problem" + e.Error())
		return
	}

	_ = console.WriteLine("Result:%s", result)
}

func Run() {
	console.CheckLengthEquals(3, len(os.Args), "./rpc_client <host> <port>")
	client := connect()

	defer func(c *rpc.Client) {
		_ = c.Close()
	}(client)

	for {
		count := console.ReadInt("Input count:", "Invalid value!...")

		if count <= 0 {
			break
		}

		callGenerate(client, &count)
	}
}
