package main

import(
    "fmt"
    "context"
    pb "aplum.com/helloworld"
    "net"
    "google.golang.org/grpc"
)

const(
    port = ":8080"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message : "Hello" + in.Name}, nil
}

func main() {
    // 指定执行程序监听的端口
    lis, err := net.Listen("tcp", port)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // 建立gRPC服务器，并注册服务
    s := grpc.NewServer()
    pb.RegisterGreetServer(s, &server)
    fmt.Println("Server run ...")

    // 启动服务
    if err := s.Server(lis); err != nil {
        fmt.Println(err.Error())
    }
}
