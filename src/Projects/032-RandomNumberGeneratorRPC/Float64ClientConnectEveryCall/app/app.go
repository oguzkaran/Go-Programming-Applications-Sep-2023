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

func callGenerate(client *rpc.Client, info *shared.RandomNumberInfo[float64]) {
	var result float64
	e := client.Call("RandomNumberGenerator.GenerateNumberFloat64", info, &result)
	if e != nil {
		_, _ = os.Stderr.WriteString("RPC problem" + e.Error())
		return
	}

	_ = console.WriteLine("Result:%f", result)
}

func Run() {
	console.CheckLengthEquals(3, len(os.Args), "./rpc_client <host> <port>")

	for {
		minVal := console.ReadFloat64("Input origin:", "Invalid value!...")
		boundVal := console.ReadFloat64("Input bound:", "Invalid value!...")

		if minVal == 0 && boundVal == 0 {
			break
		}

		client := connect()

		defer func(c *rpc.Client) {
			_ = c.Close()
		}(client)

		info := &shared.RandomNumberInfo[float64]{Min: minVal, Bound: boundVal}
		callGenerate(client, info)
	}
}
