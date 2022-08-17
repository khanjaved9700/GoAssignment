package main

import (
	"log"
	"net"

	"wc_grpc/implementation"
	pb "wc_grpc/protoFiles"

	"google.golang.org/grpc"
)

const port = ":50051"

//main function
func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen0: ", err)
	}

	//Initialize new Server
	s := grpc.NewServer()

	//Regester the server as a new grpc service
	pb.RegisterWordCountServiceServer(s, &implementation.WcServer{})

	// listenning PORT 50051 at this address
	log.Println("server listening at ", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}

}
