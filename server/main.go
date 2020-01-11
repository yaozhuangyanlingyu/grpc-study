package main

import(
	"net"
	"log"
	"context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
)

const(
	port = ":8661"
)

// 服务对象
type server struct{}

// SayHello 实现服务的接口 在proto中定义所有服务都是接口
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", port)
	}

	// 起一个服务
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	// 注册反射服务 这个服务是CLI使用的 跟服务本身没关系
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}