package main

import (
	"bookms/grpc/pb"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = ":40021"

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func HandleError(err error) {
	log.Fatal(err)
}

func main() {

	// dial a connection to the grpc server
	// conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("dial connection failed...%v", err)
	}
	defer conn.Close()
	// create new client
	client := pb.NewBookServiceClient(conn)

	// context initialization
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	//  take input from console....
	// made a menu bar like in a resturent menubar where custumer give order....

	for {
		fmt.Println(`*********************MENU*********************
		*	1. UPLOAD A NEW BOOK	*
		*	2. GET ALL BOOK			*
		*	3. SEARCH BOOK			*
		*	4. UPDATE				*
		*	5. DELETE				*
		for exit choose any ohter number...
		`)
		fmt.Println("_________________________")
		fmt.Println("choose options....")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("wrong input your exit now!", err)
		}
		option, err := strconv.ParseInt(strings.TrimSpace(input), 0, 0)
		if err != nil {
			log.Fatal("Error while parsing", err)

		}

		switch option {
		case 1:
			err := BookCreate(client, ctx)
			HandleError(err)
			continue

		case 2:
			err := BookGetAll(client, ctx)
			HandleError(err)
			continue
		case 3:
			err := BookSearch(client, ctx)
			HandleError(err)
			continue
		case 4:
			err := BookUpdate(client, ctx)
			HandleError(err)
			continue
		case 5:
			err := BookDelete(client, ctx)
			HandleError(err)
			continue
		default:
			os.Exit(0)
		}

	}
}

// func BookDelete(client pb.BookServiceClient, ctx context.Context) {
// 	panic("unimplemented")
// }

// func BookUpdate(client pb.BookServiceClient, ctx context.Context) {
// 	panic("unimplemented")
// }
