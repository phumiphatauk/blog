package gapi

import (
	db "blog_api/db/sqlc"
	"blog_api/pb"
	"blog_api/token"
	"blog_api/util"
	"blog_api/worker"
	"fmt"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedBlogServiceServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type jsonResponseWithPaginate struct {
	jsonResponse
	Total int64 `json:"total"`
}

// NewServer creates a new gRPC server and set up routing.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
