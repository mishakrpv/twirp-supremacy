package impl

import (
	"context"
	"strings"

	pb "github.com/mishakrpv/twirp-supremacy/rpc/server"
)

type Server struct{}

var a pb.Server = &Server{}

func (s *Server) Yell(ctx context.Context, req *pb.Material) (*pb.Shout, error) {
	msg := strings.ToUpper(req.Content)
	shoutContent := make([]string, req.Count)

	for range req.Count {
		shoutContent = append(shoutContent, msg)
	}

	lamentation := "gosh"

	return &pb.Shout{
		Content:     shoutContent,
		Lamentation: &lamentation,
	}, nil
}
