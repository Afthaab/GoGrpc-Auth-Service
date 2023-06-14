package api

import (
	"log"
	"net"

	handler "github.com/auth/service/pkg/api/handler"
	pb "github.com/auth/service/pkg/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ServerHttp struct {
	Engine *gin.Engine
}

func NewGrpcServer(userHandler *handler.UserHandler, grpcPort string) {
	// create a listner server for grpc
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalln("Failed to listen to the GRPC Port", err)
	}
	//create a new grpc server
	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Could not serve the GRPC Server: ", err)
	}

}

func NewServerHttp(userHandler *handler.UserHandler) *ServerHttp {
	engine := gin.New()

	//call grpc func
	go NewGrpcServer(userHandler, "8889")

	engine.Use(gin.Logger())

	return &ServerHttp{
		Engine: engine,
	}
}

func (ser *ServerHttp) Start() {
	ser.Engine.Run(":7777")
}
