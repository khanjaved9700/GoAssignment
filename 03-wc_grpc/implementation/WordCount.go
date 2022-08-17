package implementation

import (
	"context"
	"errors"
	"log"
	"regexp"
	"sort"
	"strings"
	pb "wc_grpc/protoFiles"
)

type WcServer struct {
	pb.UnimplementedWordCountServiceServer
}

//WordCount Function
func (s *WcServer) WordCount(ctx context.Context, in *pb.Request) (*pb.Response, error) {

	log.Println("Recieved: ", in.Text)
	content := strings.TrimSpace(in.Text)
	//Checking empty string
	if content == "" {
		return nil, errors.New("empty string")

	}

	//Removing all spcial charecters and white spaces from string
	reg, _ := regexp.Compile(`[^\w]`)

	content = reg.ReplaceAllString(content, " ")

	//making slice of a content
	strSlice := strings.Split(content, " ")

	//making dictionarry and store count in it
	wcMap := make(map[string]int32)

	for _, word := range strSlice {
		if word != "" {

			wcMap[word]++ // increment here

		}

	}

	//Check map empty!!!
	if len(wcMap) == 0 {
		return nil, errors.New("not a valid word")
	}

	//A slice that contains words
	words := make([]string, 0, len(wcMap))
	for w := range wcMap {
		words = append(words, w)
	}

	//Sorting Words by count in slice
	sort.Slice(words, func(i, j int) bool {
		return wcMap[words[i]] > wcMap[words[j]]
	})

	var wcList []*pb.WordCount
	for key, value := range wcMap {
		wcList = append(wcList, &pb.WordCount{Word: key, Count: value})
	}

	sort.Slice(wcList, func(i, j int) bool {
		return wcList[i].Count > wcList[j].Count
	})

	return &pb.Response{WcList: wcList}, nil
}
