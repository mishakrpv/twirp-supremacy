package main

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/mishakrpv/twirp-supremacy/rpc/server"
)

func main() {
	client := pb.NewServerProtobufClient("http://localhost:8080", &http.Client{})

	resp, err := client.Yell(context.Background(), &pb.Material{Content: "I use arch btw", Count: 2})
	if err == nil {
		for _, msg := range resp.Content {
			fmt.Println(msg)
		}
		if resp.Lamentation != nil {

			fmt.Println(resp.Lamentation)
		}
	}
}
