package password

import (
	"Client/generator/random/password/shared"
	"net/rpc"
)

type ConnectionInfo struct {
	client *rpc.Client
}

func (info *ConnectionInfo) Close() error {
	return info.client.Close()
}

func callGenerate(client *rpc.Client, info *shared.PasswordInfo) (*shared.PasswordsInfo, error) {
	var result shared.PasswordsInfo

	e := client.Call("RandomPasswordGenerator.GeneratePasswords", info, &result)
	if e != nil {
		return nil, e
	}

	return &result, nil
}

func Connect(address string) (*ConnectionInfo, error) {
	c, e := rpc.Dial("tcp", address)
	if e != nil {
		return nil, e
	}

	return &ConnectionInfo{client: c}, nil
}

func GeneratePasswords(info *ConnectionInfo, origin, bound, count int) ([]string, error) {
	passWordInfo := &shared.PasswordInfo{Range: shared.RandomRange{Origin: origin, Bound: bound}, Count: count}
	r, e := callGenerate(info.client, passWordInfo)

	if e != nil {
		return nil, e
	}

	return r.Passwords, e
}

func GeneratePasswordsConnect(address string, origin, bound, count int) ([]string, error) {
	info, e := Connect(address)
	if e != nil {
		return nil, e
	}

	defer func(info *ConnectionInfo) {
		_ = info.Close()
	}(info)

	return GeneratePasswords(info, origin, bound, count)
}
