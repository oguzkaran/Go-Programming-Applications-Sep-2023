package app

import (
	"Client/csd/console"
	"Client/shared"
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

func callGenerate(client *rpc.Client, r *shared.RandomRange) {
	var result string

	e := client.Call("RandomPasswordGenerator.GeneratePasswordRandomRange", r, &result)
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
		origin := console.ReadInt("Input origin:", "Invalid value!...")
		bound := console.ReadInt("Input bound:", "Invalid value!...")

		if origin == 0 && bound == 0 {
			break
		}
		r := &shared.RandomRange{Origin: origin, Bound: bound}
		callGenerate(client, r)
	}
}
