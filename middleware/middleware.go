package middleware

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func LoggingInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (
	resp any,
	err error,
) {
	startTime := time.Now()

	// Call the handler (the actual RPC method)
	resp, err = handler(ctx, req)

	// Calculate the duration of the RPC call
	duration := time.Since(startTime)

	// Log the method, duration, and error status
	fmt.Printf("Method: %s, Duration: %v, Error: %v\n", info.FullMethod, duration, err)

	return resp, err
}
