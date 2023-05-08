package grpc_stuff

import (
	"context"
)

var (
	_ ServeCacheServer = (*Server)(nil)
)

type Server struct {
}

func (s *Server) Set(ctx context.Context, insertRequest *SetRequest) (*Empty, error) {
	return &Empty{}, nil
}

func (s *Server) SetDefault(ctx context.Context, setDefaultRequest *SetDefaultRequest) (*Empty, error) {
	return &Empty{}, nil
}

func (s *Server) Add(ctx context.Context, addRequest *AddRequest) (*Empty, error) {
	return &Empty{}, nil
}

func (s *Server) Replace(ctx context.Context, replaceRequest *ReplaceRequest) (*Empty, error) {
	return &Empty{}, nil
}

func (s *Server) Get(ctx context.Context, getRequest *GetRequest) (*GetResponse, error) {
	return &GetResponse{}, nil
}

func (s *Server) GetWithExpiration(ctx context.Context, getWithExpirationRequest *GetWithExpirationRequest) (*GetWithExpirationResponse, error) {
	return &GetWithExpirationResponse{}, nil
}

func (s *Server) Delete(ctx context.Context, deleteRequest *DeleteRequest) (*Empty, error) {
	return &Empty{}, nil
}

func (s *Server) DeleteExpired(ctx context.Context, _ *Empty) (*Empty, error) {
	return &Empty{}, nil
}

func (s *Server) Items(ctx context.Context, _ *Empty) (*ItemsResponse, error) {
	return &ItemsResponse{}, nil
}

func (s *Server) ItemCount(ctx context.Context) {
}

func (s *Server) Flush() {}

// stub
func (s *Server) mustEmbedUnimplementedServeCacheServer() {}
