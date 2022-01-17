package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func GetTokenByName(ctx context.Context, header string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "No tokens in metadata.")
	}

	value := md.Get(header)
	if len(value) > 0 {
		return value[0], nil
	}

	return "", status.Error(codes.Unauthenticated, header+" is empty.")
}
