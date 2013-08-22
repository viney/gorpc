package user

import (
	"log"
	"net/rpc"
	"sync"
)

const (
	ClientName = "user"
	ClientUri  = "127.0.0.1:8080"
)

var mutex sync.Mutex

type UserArg struct {
	Uid int
}

type UserRet struct {
	Name  string
	Email string
}

func Query(uid int) (name, email string, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	client, err := dial()
	if err != nil {
		return "", "", err
	}
	defer client.Close()

	in := &UserArg{uid}
	ret := &UserRet{}
	if err := client.Call(ClientName+".Query", in, ret); err != nil {
		log.Fatal("client.Call: ", err)
		return "", "", err
	}

	return ret.Name, ret.Email, nil
}

func dial() (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", ClientUri)
	if err != nil {
		return nil, err
	}

	return client, nil
}
