package app

import (
	"Client/csd/console"
	"Client/generator/random/password"
	"os"
)

func Run() {
	console.CheckLengthEquals(3, len(os.Args), "./rpc_client <host> <port>")
	c, e := password.Connect(os.Args[1] + ":" + os.Args[2])

	if e != nil {
		_, _ = os.Stderr.WriteString("Connection failed!...")
		os.Exit(1)
	}

	defer func(c *password.ConnectionInfo) {
		_ = c.Close()
	}(c)

	for {
		origin := console.ReadInt("Input origin:", "Invalid value!...")
		bound := console.ReadInt("Input bound:", "Invalid value!...")
		count := console.ReadInt("Input count:", "Invalid value!...")

		if origin == 0 && bound == 0 && count <= 0 {
			break
		}
		p, e := password.GeneratePasswords(c, origin, bound, count)
		if e != nil {
			_, _ = os.Stderr.WriteString(e.Error())
		} else {
			_ = console.WriteLine("%v", p)
		}
	}
}
