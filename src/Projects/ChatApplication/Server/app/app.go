package app

import (
	"Server/csd/cerror"
	"Server/csd/console"
	"Server/csd/net/tcp"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func registerServerCallback(socket net.Conn) {
	defer socket.Close()
	defer wg.Done()
	//REGISTER <name> <nickname> <password> <accept> <accept or reject>
	//If nickname not exists save to users database
}

func loginServerCallback(socket net.Conn) {
	defer socket.Close()
	defer wg.Done()
	//LOGIN <nickname> <password>
	//If nickname and password is correct
	//save the login datetime to database and ...
	//If nickname exists but password is wrong
	//ERROR_PASSWORD_MISMATCH
	//else if nickname does not exist
	//ERROR_NICKNAME_NOT_FOUND send to client
	//client send if use as a guest or not
	//If client will be guest theb ...
}

func Run() {
	console.CheckLengthEquals(2, len(os.Args), "usage: ./server <base port>")
	basePort, err := strconv.Atoi(os.Args[1])
	cerror.CheckError(err, "Invalid base port number")

	registerAddress := fmt.Sprintf(":%d", basePort)
	loginAddress := fmt.Sprintf(":%d", basePort+1)

	wg.Add(2)
	go tcp.StartServer(registerAddress, "RegisterServer:\n", "", registerServerCallback)
	go tcp.StartServer(loginAddress, "LoginServer:\n", "", loginServerCallback)
	wg.Wait()
}
