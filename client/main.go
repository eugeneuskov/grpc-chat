package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/eugeneuskov/grpc-chat/proto/pb"
	"google.golang.org/grpc"
	"log"
	"os"
	"sync"
	"time"
)

var client pb.BroadcastClient
var wg *sync.WaitGroup

func init() {
	wg = &sync.WaitGroup{}
}

func connect(user *pb.User) error {
	var streamError error

	stream, err := client.Connect(context.Background(), &pb.ConnectRequest{
		User:   user,
		Active: true,
	})
	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}

	wg.Add(1)
	go func(stream pb.Broadcast_ConnectClient) {
		defer wg.Done()

		for {
			msg, err := stream.Recv()
			if err != nil {
				streamError = fmt.Errorf("error reading message: %v", err)
				break
			}

			fmt.Printf("%v : %s\n", msg.GetUser().GetName(), msg.GetMessage())
		}
	}(stream)

	return streamError
}

func main() {
	timestamp := time.Now()
	done := make(chan int)

	name := flag.String("N", "Anon", "The name of the user")
	flag.Parse()

	id := sha256.Sum256([]byte(timestamp.String() + *name))

	conn, err := grpc.Dial("localhost:11011", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldnt connect to server: %s\n", err.Error())
	}

	client = pb.NewBroadcastClient(conn)
	user := &pb.User{
		Id:   hex.EncodeToString(id[:]),
		Name: *name,
	}

	_ = connect(user)

	wg.Add(1)
	go func() {
		defer wg.Done()

		scanner := bufio.NewScanner(os.Stdin)
		timestamp := time.Now().String()
		messageId := sha256.Sum256([]byte(timestamp + *name))

		for scanner.Scan() {
			_, err := client.SendMessage(context.Background(), &pb.Content{
				Id:        hex.EncodeToString(messageId[:]),
				User:      user,
				Message:   scanner.Text(),
				Timestamp: timestamp,
			})
			if err != nil {
				fmt.Printf("error sendong message: %s\n", err.Error())
				break
			}
		}
	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done
}
