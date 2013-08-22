package main

import (
	"fmt"
	"net"
	"net/rpc"

	"engine/user"
)

type UserServer interface {
	Query(*user.UserArg, *user.UserRet) error
}

func UserService(clientName, clientUri string) (run func(), err error) {
	service := NewUserServer()

	server := rpc.NewServer()
	if err := server.RegisterName(clientName, service); err != nil {
		return nil, err
	}

	run = func() {
		l, err := net.Listen("tcp", clientUri)
		if err != nil {
			fmt.Println("net.Listen: ", err)
		}
		server.Accept(l)
	}

	return run, nil
}

type userServer struct {
	uid int
}

func NewUserServer() UserServer {
	return &userServer{}
}

func (u *userServer) Query(in *user.UserArg, ret *user.UserRet) error {
	var users = map[int][]string{
		1: []string{"viney", "viney.chow@gmail.com"},
	}

	uid := in.Uid
	if v, ok := users[uid]; ok {
		ret.Name = v[0]
		ret.Email = v[1]
	} else {
		ret.Name = "Null"
		ret.Email = "Null"
	}

	return nil
}
