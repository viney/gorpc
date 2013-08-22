package main

import (
	"fmt"

	"engine/signal"
	"engine/user"
)

func main() {
	fmt.Println("user service start...")
	run, err := UserService(user.ClientName, user.ClientUri)
	if err != nil {
		fmt.Println("UserServer: ", err)
		return
	}

	go run()
	signal.Serve()
	fmt.Println("user service end...")
}
