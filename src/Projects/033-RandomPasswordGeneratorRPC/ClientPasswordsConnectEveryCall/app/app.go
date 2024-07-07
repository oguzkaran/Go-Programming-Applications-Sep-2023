package app

import (
	"Client/csd/console"
	"Client/generator/random/password"
	"os"
)

func Run() {
	console.CheckLengthEquals(3, len(os.Args), "./rpc_client <host> <port>")
	address := os.Args[1] + ":" + os.Args[2]

	for {
		origin := console.ReadInt("Input origin:", "Invalid value!...")
		bound := console.ReadInt("Input bound:", "Invalid value!...")
		count := console.ReadInt("Input count:", "Invalid value!...")

		if origin == 0 && bound == 0 && count <= 0 {
			break
		}
		p, e := password.GeneratePasswordsConnect(address, origin, bound, count)
		if e != nil {
			_, _ = os.Stderr.WriteString(e.Error())
		} else {
			_ = console.WriteLine("%v", p)
		}
	}
}
