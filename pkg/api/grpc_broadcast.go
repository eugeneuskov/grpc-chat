package api

import (
	"context"
	"github.com/eugeneuskov/grpc-chat/pkg/services"
	"github.com/eugeneuskov/grpc-chat/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"sync"
	"time"
)

type connection struct {
	stream pb.Broadcast_ConnectServer
	id     string
	active bool
	error  chan error
}

type grpcBroadcastService struct {
	service     services.Auth
	connections []*connection
}

func (b *grpcBroadcastService) Connect(request *pb.ConnectRequest, server pb.Broadcast_ConnectServer) error {
	conn := &connection{
		stream: server,
		id:     request.GetUser().GetId(),
		active: true,
		error:  make(chan error),
	}

	b.connections = append(b.connections, conn)

	b.send(&pb.Content{
		Id:        "",
		User:      request.GetUser(),
		Message:   "",
		Timestamp: time.Now().String(),
		Type:      "chat",
	})

	return <-conn.error
}

func (b *grpcBroadcastService) SendMessage(_ context.Context, content *pb.Content) (*emptypb.Empty, error) {
	b.send(content)
	return &emptypb.Empty{}, nil
}

func (b *grpcBroadcastService) send(content *pb.Content) {
	wg := sync.WaitGroup{}
	done := make(chan int)

	for _, conn := range b.connections {
		wg.Add(1)

		go func(msg *pb.Content, conn *connection) {
			defer wg.Done()

			if conn.active {
				if err := conn.stream.Send(msg); err != nil {
					conn.active = false
					conn.error <- err
				}
			}
		}(content, conn)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done
}

func newGrpcBroadcastService(s *grpc.Server, service services.Auth) {
	if s != nil {
		pb.RegisterBroadcastServer(s, &grpcBroadcastService{service: service})
	}
}
