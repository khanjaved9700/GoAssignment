package main

import (
	pb "bookms/grpc/pb"
	c "bookms/pkg/lms"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":55001"

func main() {
	fmt.Println("Welocm to Library Management System")

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Server connection Failed", err)
	}

	// server Installation
	server := grpc.NewServer()

	// registering server at new grpc server....

	pb.RegisterBookServiceServer(server, &c.BookServiceServer{})

	log.Println("server listening at", listen.Addr())

	if err := server.Serve(listen); err != nil {
		log.Fatal("Failed", err)
	}
}
