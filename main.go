package main

import (
	user "demo/kitex_gen/user/login"
	"log"
)

func main() {
	svr := user.NewServer(new(LoginImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
