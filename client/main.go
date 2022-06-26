package main

import (
	"context"
	"demo/kitex_gen/user"
	"demo/kitex_gen/user/login"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func main() {
	cli, err := login.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &user.LoginRequest{Username: "fwh", Password: "sy"}
	resp, err := cli.CheckUser(context.Background(), req, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
