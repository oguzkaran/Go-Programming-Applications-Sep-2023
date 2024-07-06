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

func callGenerate(client *rpc.Client, info *shared.PasswordInfo) {
	var result shared.PasswordsInfo

	e := client.Call("RandomPasswordGenerator.GeneratePasswords", info, &result)
	if e != nil {
		_, _ = os.Stderr.WriteString("RPC problem" + e.Error())
		return
	}

	_ = console.WriteLine("Result:%v", result.Passwords)
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
		count := console.ReadInt("Input count:", "Invalid value!...")

		if origin == 0 && bound == 0 && count <= 0 {
			break
		}
		info := &shared.PasswordInfo{Range: shared.RandomRange{Origin: origin, Bound: bound}, Count: count}
		callGenerate(client, info)
	}
}
