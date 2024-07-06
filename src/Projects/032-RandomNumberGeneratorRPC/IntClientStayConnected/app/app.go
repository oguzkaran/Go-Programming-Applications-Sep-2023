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

func callGenerate(client *rpc.Client, info *shared.RandomNumberInfo[int]) {
	var result int
	e := client.Call("RandomNumberGenerator.GenerateNumberInt", info, &result)
	if e != nil {
		_, _ = os.Stderr.WriteString("RPC problem" + e.Error())
		return
	}

	_ = console.WriteLine("Result:%d", result)
}

func Run() {
	console.CheckLengthEquals(3, len(os.Args), "./rpc_client <host> <port>")
	client := connect()

	defer func(c *rpc.Client) {
		_ = c.Close()
	}(client)

	for {
		minVal := console.ReadInt("Input origin:", "Invalid value!...")
		boundVal := console.ReadInt("Input bound:", "Invalid value!...")

		if minVal == 0 && boundVal == 0 {
			break
		}
		info := &shared.RandomNumberInfo[int]{Min: minVal, Bound: boundVal}
		callGenerate(client, info)
	}
}
