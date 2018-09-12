package api

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"time"
	"github.com/blog/consul"
	"github.com/blog/handler"
	pb "github.com/blog/proto"
)

type BlogService struct{}

func (*BlogService) Start() {
	lis, err := net.Listen("tcp", "127.0.0.1:9100")
	if err != nil {
		log.Fatal("BlogService: " + err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &BlogService{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("BlogService: " + err.Error())
	}
}

func init() {
	consul.Register("BlogService", "127.0.0.1", 9100, "127.0.0.1:8500", 10*time.Second, 15)
	handler.RegisterHTTPHandler("BlogService", pb.RegisterBlogServiceHandler)
}
