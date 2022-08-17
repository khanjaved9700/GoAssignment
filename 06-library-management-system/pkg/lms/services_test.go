package lms

import (
	pb "bookms/grpc/pb"
	// c "bookms/pkg/lms"
	"context"
	"log"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {

	lis = bufconn.Listen(bufSize)
	server := grpc.NewServer()
	pb.RegisterBookServiceServer(server, &BookServiceServer{})
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

// func bufDialer(context.Context, string) (net.Conn, error) {
// 	return lis.Dial()
// }

func TestService(t *testing.T) {

	//Dial a connection to grpc Server
	ctx := context.Background()
	conn, err := grpc.Dial(":97001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	//Create new Client
	client := pb.NewBookServiceClient(conn)

	//Test CreateBook
	newBook := &pb.Book{
		Title:  "test title",
		Author: "test author",
	}

	response, err := client.CreateBook(ctx, &pb.CreateBookRequest{Book: newBook})
	if err != nil {
		t.Error("Test CreateBook FAILED!\nerr: ", err)
	}
	t.Log("Test CreateBook PASSED.")

	//get id
	id := response.Book.Id

	_, err = client.GetAllBook(ctx, &pb.GetAllBookRequest{})
	if err != nil {
		t.Error("Test ListAllBooks FAILED!\nerr: ", err)
	}
	t.Log("Test ListAllBooks PASSED.")

	//Test SearchBooks
	_, err = client.SearchBook(ctx, &pb.SearchBookRequest{
		SearchBy: &pb.SearchBookRequest_Title{Title: newBook.Title},
	})
	if err != nil {
		t.Error("Test SearchBooks FAILED!\nerr: ", err)
	}

	t.Log("Test SearchBooks PASSED.")

	//Test Updatebook
	updateBook := &pb.Book{
		Id:     id,
		Title:  "updated test title",
		Author: "updated test title",
	}
	_, err = client.UpdateBook(ctx, &pb.UpdateBookRequest{Book: updateBook})
	if err != nil {
		t.Error("Test UpdateBooks FAILED!\nerr: ", err)
	}

	t.Log("Test UpdateBooks PASSED.")

	//Test DeleteBook
	_, err = client.DeleteBook(ctx, &pb.DeleteBookRequest{})
	if err != nil {
		t.Error("Test DeleteBooks FAILED!\nerr: ", err)
	}

	t.Log("Test DeleteBooks PASSED.")

}
