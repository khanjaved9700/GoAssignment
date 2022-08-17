package lms

import (
	pb "bookms/grpc/pb"
	"bookms/pkg/model"
	"context"
	"errors"
	"log"

	"bookms/pkg/database"
)

//Implemetation of proto(grpc.pb.go) interfaces
type BookServiceServer struct {
	pb.UnimplementedBookServiceServer
}

//Create Book
func (s *BookServiceServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	book := req.Book

	//request Validation
	if book.Title == "" || book.Author == "" {
		return nil, errors.New("invalid Data")
	}

	bookData := model.Book{
		Title:  book.Title,
		Author: book.Author,
	}

	id, err := database.CreateBook(ctx, bookData)

	if err != nil {
		return nil, err
	}

	book.Id = uint64(id)

	return &pb.CreateBookResponse{Book: &pb.Book{Id: book.Id, Title: book.Title, Author: book.Author}}, nil
}

//Get All Books
func (s *BookServiceServer) GetAllBook(req *pb.GetAllBookRequest, stream pb.BookService_GetAllBookServer) error {

	list, err := database.GetAllBooks()
	if err != nil {
		return err

	}

	for _, book := range list {
		err = stream.Send(&pb.GetAllBookResponse{Book: &pb.Book{Id: uint64(book.ID), Title: book.Title, Author: book.Author}})

		if err != nil {
			return err
		}
	}

	return nil
}

//Search Book
func (s *BookServiceServer) SearchBook(req *pb.SearchBookRequest, stream pb.BookService_SearchBookServer) error {

	title := req.GetTitle()
	author := req.GetAuthor()

	if title == "" && author == "" {
		return errors.New("nothing to search, empty argment")
	}

	search, err := database.SearchBook(title, author)

	if err != nil {
		return err
	}

	for _, book := range search {
		err = stream.Send(&pb.SearchBookResponse{Book: &pb.Book{Id: uint64(book.ID), Title: book.Title, Author: book.Author}})

		if err != nil {
			return err
		}
	}

	return nil
}

//Update Book

func (s *BookServiceServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	bookUpdates := req.GetBook()

	if bookUpdates.Title == "" || bookUpdates.Author == "" || bookUpdates.Id == 0 {
		log.Fatal("Invalid Data")
	}

	bookData := model.Book{
		Title:  bookUpdates.Title,
		Author: bookUpdates.Author,
	}
	id, err := database.UpdateBook(ctx, bookData, bookUpdates.Id)
	if err != nil {
		log.Fatal(err)
	}
	bookUpdates.Id = uint64(id)

	return &pb.UpdateBookResponse{Book: &pb.Book{Id: bookUpdates.Id, Title: bookData.Title, Author: bookData.Author}}, nil
}

//Delete

func (s *BookServiceServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	//Validation
	del := req.GetTitle()
	if del == "" {
		return nil, errors.New("invalid Delete Request")
	}

	err := database.DeleteBook(del)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteBookResponse{Dlt: true}, nil
}
