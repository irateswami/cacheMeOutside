package grpc_stuff

import "context"

var _ ServeCacheServer = (*Server)(nil)

type Server struct{}

func (s *Server) Get(ctx context.Context, getRequest *GetRequest) (*GetResponse, error) {
	return &GetResponse{}, nil
}

func (s *Server) Insert(ctx context.Context, insertRequest *InsertRequest) (*Empty, error) {
	return &Empty{}, nil
}

func (s *Server) mustEmbedUnimplementedServeCacheServer() {}
