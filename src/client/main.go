package main

import (
	"flag"
	"log"
	"strconv"

	"engine/user"
)

func main() {
	flag.Parse()
	uid, err := strconv.Atoi((flag.Arg(0)))
	if err != nil {
		log.Fatal("strconv.Atoi: ", err)
		return
	}

	name, email, err := user.Query(uid)
	if err != nil {
		log.Fatal("user.Query: ", err)
		return
	}

	log.Println(name, email)
}
