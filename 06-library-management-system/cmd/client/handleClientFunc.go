package main

import (
	"bookms/grpc/pb"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

// creating a book

func BookCreate(client pb.BookServiceClient, ctx context.Context) error {
	// provide input

	fmt.Println("\nPlease Enter Details:")
	fmt.Println("\nBook Titles: ")

	title, err := reader.ReadString('\n')
	if err != nil {
		// return errors.New(fmt.Sprint("Invalid Input: Please try again", err))
		HandleError(err)
	}
	title = strings.TrimSpace(title)

	// fmt.Println(title)

	fmt.Println("\nAuthor Name: ")

	author, err := reader.ReadString('\n')
	if err != nil {
		// return errors.New(fmt.Sprint("Incorrect! please provide correct author name"))
		HandleError(err)
	}
	author = strings.TrimSpace(author)

	// fmt.Println(author)

	if title == "" || author == "" {
		// return errors.New("empty strings plz do some enteries...")
		fmt.Println("Please check Title or Author It coudn't be empty")
	}
	// creating book to send request...
	NewBook := &pb.Book{
		Title:  title,
		Author: author,
	}

	// call createbook that returns a responsefunc

	res, err := client.CreateBook(ctx, &pb.CreateBookRequest{Book: NewBook})

	if err != nil {
		// return errors.New(fmt.Sprint("book response failed...", err))
		HandleError(err)
	}

	// print new uploaded book....

	log.Printf(`New Book Uploaded
	BOOK ID:%d
	TITLE: %s
	AUTHOR: %s`, res.Book.Id, res.Book.Title, res.Book.Author)

	return nil
}

//  Getting all Books

func BookGetAll(client pb.BookServiceClient, ctx context.Context) error {

	fmt.Println("Enter a page No. :")
	page, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// HandleError(err)
	page = strings.TrimSpace(page)
	pageNo, err := strconv.ParseInt(page, 0, 0)
	if err != nil {
		log.Fatal("Error While Parsing")
	}

	fmt.Println("Page size...")
	pn, _ := reader.ReadString('\n')
	pn = strings.TrimSpace(pn)
	ps, err := strconv.ParseInt(pn, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	print := (pageNo - 1) * ps

	// call GetAllBooks thats returns a stream....
	stream, err := client.GetAllBook(ctx, &pb.GetAllBookRequest{})
	if err != nil {
		// return errors.New(fmt.Sprint("streaming failed..."))
		HandleError(err)
	}
	fmt.Println("Books: ")

	for i := 0; i <= int(print); i++ {
		// stream.Recv() function returns a pointer to the book in a current iteration
		responce, err := stream.Recv()
		// if end of stream then break the loop
		if err == io.EOF {
			break
		}
		if err != nil {
			// return errors.New(fmt.Sprint("Streaming  failed"))
			HandleError(err)
		}
		if i >= int(print) {
			fmt.Printf("%d,%v\n", i+1, responce.GetBook())
		}
	}
	return nil
}

// Book Searching...... Bi-Directional Streaming
func BookSearch(client pb.BookServiceClient, ctx context.Context) error {

	var stream pb.BookService_SearchBookClient

	fmt.Println(`Choose an option...
	
	1. Search by Book Title
	2. Search by Book Author
	
	Choose: `)

	input, err := reader.ReadString('\n')
	if err != nil {
		// return err
		HandleError(err)
	}

	input = strings.TrimSpace(input)
	Choice, err := strconv.ParseInt(input, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing...")
	}
	switch Choice {
	case 1:
		fmt.Println("Enter Title: ")
		title, err := reader.ReadString('\n')
		if err != nil {
			// return err
			HandleError(err)
		}

		title = strings.TrimSpace(title)

		if title == "" {
			// return errors.New("empty title")
			HandleError(err)
		}

		if err != nil {
			// return err
			HandleError(err)
		}

		fmt.Println("Books we have found...")

		// Start iterating
		// for i := 0; i < 3; i++ {

		//Streaming
		stream, err = client.SearchBook(ctx, &pb.SearchBookRequest{
			SearchBy: &pb.SearchBookRequest_Title{Title: title},
		})

		if err != nil {
			// return err
			// HandleError(err)
			fmt.Println("Error while streaming...")
		}
		// }

	case 2:
		fmt.Println("Enter Book Author: ")
		author, err := reader.ReadString('\n')
		if err != nil {
			// return err
			HandleError(err)

		}

		author = strings.TrimSpace(author)

		if author == "" {
			// return errors.New("empty author")
			HandleError(err)
		}

		//Streaming
		stream, err = client.SearchBook(ctx, &pb.SearchBookRequest{
			SearchBy: &pb.SearchBookRequest_Author{Author: author},
		})

		if err != nil {
			// return err
			HandleError(err)
		}

	default:
		return errors.New("search book failed")
		// fmt.Println("Please Choose Right Selection")
	}

	fmt.Println("Books we have found...")

	// Start iterating
	for i := 0; i < 3; i++ {

		// stream.Recv returns a pointer to a book in a current iteration
		responseStream, err := stream.Recv()
		// If end of stream, break the loop
		if err == io.EOF {
			break
		}
		// if err, print error
		if err != nil {
			// return errors.New(fmt.Sprint("Stream error: ", err))
			HandleError(err)
		}

		// If everything went well use the generated getter to print the Book Details
		fmt.Println(responseStream.GetBook())

	}

	return nil
}

//Update Book
func BookUpdate(client pb.BookServiceClient, ctx context.Context) error {

	fmt.Println("Updating book by id Please Enter Book ID:")
	id, err := reader.ReadString('\n')
	if err != nil {
		// return errors.New(fmt.Sprint("Invalid ID..", err))
		HandleError(err)
	}
	id = strings.TrimSpace(id)

	fmt.Print("Book Title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		// return errors.New(fmt.Sprint("Invalid Book Title..", err))
		HandleError(err)
	}
	title = strings.TrimSpace(title)

	fmt.Print("Book Author: ")
	author, err := reader.ReadString('\n')
	if err != nil {
		// return errors.New(fmt.Sprint("Invalid Book Author..", err))
		HandleError(err)
	}
	author = strings.TrimSpace(author)

	if id == "" || title == "" || author == "" {
		// return errors.New(fmt.Sprint("Empty..."))
		HandleError(err)
	}

	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// log.Fatal(err)
		HandleError(err)
	}
	//Creating Request
	updateBook := &pb.Book{
		Id:     uint64(num),
		Title:  title,
		Author: author,
	}
	//Call UpdateBook that returns a Book as response
	response, err := client.UpdateBook(ctx, &pb.UpdateBookRequest{Book: updateBook})
	if err != nil {
		// return errors.New(fmt.Sprint("Could not update book: \n", err))
		HandleError(err)
	}
	//print
	log.Printf(`Book Updated:
	Book Id: %d
	Title: %s
	Author: %s`, response.Book.Id, response.Book.Title, response.Book.Author)

	return nil
}

//Delete Book
func BookDelete(client pb.BookServiceClient, ctx context.Context) error {
	// panic("unimplemented")

	// func BookDelete(client pb.BookServiceClient, ctx context.Context) error {

	fmt.Printf("Enter Book Title u want to delete:  ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	title := strings.TrimSpace(input)

	//Call DeleteBook
	_, err = client.DeleteBook(ctx, &pb.DeleteBookRequest{Title: title})
	if err != nil {
		return err
	}

	//Print Result
	fmt.Print("\nDeleted book with Book Title: ", title)
	return nil
}
