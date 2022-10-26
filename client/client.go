package main

//Initial setup https://www.freecodecamp.org/news/grpc-server-side-streaming-with-go/

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	pb "github.com/Lukski175/ChittyChat/time"
	"google.golang.org/grpc"
)

var clientName string

func main() {
	lampClock = 0

	log.Printf("Please input a name...")
	fmt.Scanln(&clientName)

	// dial server
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := pb.NewMessageStreamClient(conn)

	stream, err := client.Stream(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	stream.Send(&pb.MessageRequest{Message: clientName})

	lampClock++ //First event

	go SendLoop(stream)

	//Receive loop, just utilizing current thread
	func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				log.Fatalf("Server died")
				break
			}
			if resp.Clock > lampClock {
				lampClock = resp.Clock
			}
			messages = append(messages, resp)
			PrintChat()
		}
	}()
}

var messages []*pb.MessageReply
var lampClock int32

func SendLoop(stream pb.MessageStream_StreamClient) {
	for {
		reader := bufio.NewScanner(os.Stdin)
		reader.Scan()
		if reader.Text() != "" {
			lampClock++
			stream.Send(&pb.MessageRequest{Message: reader.Text(), Clock: lampClock})
		}
	}
}

func PrintChat() {
	//Clear console
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for _, resp := range messages {
		if resp.IsBroadcast {
			log.Printf("Server broadcast - %d: %s", resp.Clock, resp.Message)
		} else {
			log.Printf("%s - %d: %s", resp.Author, resp.Clock, resp.Message)
		}
	}
}
