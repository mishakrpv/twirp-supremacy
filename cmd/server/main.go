package main

import (
	"net/http"

	pb "github.com/mishakrpv/twirp-supremacy/rpc/server"
	"github.com/mishakrpv/twirp-supremacy/rpc/server/impl"
)

func main() {
	twirpHandler := pb.NewServerServer(&impl.Server{})

	mux := http.NewServeMux()

	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)

	http.ListenAndServe(":8080", mux)
}
