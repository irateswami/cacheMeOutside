package grpc_stuff

import (
	"context"
)

var (
	_ ServeCacheServer = (*Server)(nil)
)

type Server struct {
}

func (s *Server) Get(ctx context.Context, getRequest *GetRequest) (*GetResponse, error) {
	return &GetResponse{}, nil
}

func (s *Server) Set(ctx context.Context, insertRequest *InsertRequest) (*Empty, error) {
	return &Empty{}, nil
}

/**
func (c *cache) Set(k string, x []byte, d time.Duration) {
func (c *cache) SetDefault(k string, x []byte) {
func (c *cache) Add(k string, x []byte, d time.Duration) error {
func (c *cache) Replace(k string, x []byte, d time.Duration) error {
func (c *cache) Get(k string) ([]byte, bool) {
func (c *cache) GetWithExpiration(k string) ([]byte, time.Time, bool) {
func (c *cache) Delete(k string) {
func (c *cache) DeleteExpired() {
func (c *cache) Items() map[string]Item {
func (c *cache) ItemCount() int {
func (c *cache) Flush() {

SetDefault
Add
Replace
Get
GetWithExpiration
Delete
DeleteExpired
OnEvicted
Items
ItemCount
Flush
*/

func (s *Server) mustEmbedUnimplementedServeCacheServer() {}
