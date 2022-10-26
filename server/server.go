package main

//Initial setup https://www.freecodecamp.org/news/grpc-server-side-streaming-with-go/

import (
	"log"
	"net"
	"sync"

	pb "github.com/Lukski175/ChittyChat/time"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMessageStreamServer
}

func main() {
	lampClock = 0

	// create listiner
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterMessageStreamServer(s, server{})

	log.Println("start server")
	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

var wg sync.WaitGroup

func (s server) Stream(srv pb.MessageStream_StreamServer) error {
	wg.Add(1)
	clients++

	resp, err := srv.Recv()
	if err != nil {
		log.Fatalf("cannot receive %v", err)
	} else if resp.Clock > lampClock {
		lampClock = resp.Clock
	}
	lampClock++

	clientName := resp.Message
	message = pb.MessageReply{Message: clientName + " Joined", IsBroadcast: true}
	clientStreams[clients] = srv

	go SendMessageToClients(&message)
	go ReceiveLoop(srv, clientName)

	wg.Wait()
	return nil
}

var clientStreams = make(map[int]pb.MessageStream_StreamServer)
var clients int
var message pb.MessageReply
var lampClock int32

func ReceiveLoop(stream pb.MessageStream_StreamServer, clientName string) {
	clientID := clients
	for {
		resp, err := stream.Recv()

		if err != nil {
			lampClock++ //Disconnect is an event
			delete(clientStreams, clientID)
			message = pb.MessageReply{Message: clientName + " Disconnected", IsBroadcast: true}
			go SendMessageToClients(&message)
			break
		} else {
			if resp.Clock > lampClock {
				lampClock = resp.Clock
			}
			lampClock++
			message = pb.MessageReply{Message: resp.Message, Author: clientName}
			go SendMessageToClients(&message)
			log.Printf("Resp received: %s", resp.Message)
		}
	}
	wg.Done()
}

func SendMessageToClients(msg *pb.MessageReply) {
	lampClock++
	for _, stream := range clientStreams {
		msg.Clock = lampClock
		err := stream.Send(msg)
		if err != nil {
			log.Printf("%v", err)
		}
	}
}
