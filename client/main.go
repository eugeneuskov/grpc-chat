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
var authClient pb.AuthClient
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

			switch msg.GetType() {
			case "message":
				fmt.Printf("%s: %s\n", msg.GetUser().GetName(), msg.GetMessage())
			case "chat":
				fmt.Printf("--- %s is online\n", msg.GetUser().GetName())
			}
		}
	}(stream)

	return streamError
}

func main() {
	done := make(chan int)

	name := flag.String("N", "Anon", "The name of the user")
	password := flag.String("P", "---", "Password of the user")
	flag.Parse()

	conn, err := grpc.Dial("localhost:11011", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldnt connect to server: %s\n", err.Error())
	}
	defer conn.Close()

	/* AUTH */
	authClient = pb.NewAuthClient(conn)

	token, err := authClient.Login(context.Background(), &pb.LoginRequest{
		Login:    *name,
		Password: *password,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	userInfo, err := authClient.Info(context.Background(), &pb.InfoRequest{
		Token: token.GetToken(),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	user := &pb.User{
		Id:   userInfo.GetId(),
		Name: userInfo.GetUsername(),
	}

	client = pb.NewBroadcastClient(conn)

	if err = connect(user); err != nil {
		log.Fatal(err.Error())
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		scanner := bufio.NewScanner(os.Stdin)
		timestamp := time.Now().String()
		messageId := sha256.Sum256([]byte(timestamp + user.GetName()))

		for scanner.Scan() {
			_, err := client.SendMessage(context.Background(), &pb.Content{
				Id:        hex.EncodeToString(messageId[:]),
				User:      user,
				Message:   scanner.Text(),
				Timestamp: timestamp,
				Type:      "message",
			})
			if err != nil {
				fmt.Printf("error sending message: %s\n", err.Error())
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
