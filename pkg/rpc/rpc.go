package rpc

import (
	"fmt"
	"github.com/kaiiorg/go-bif-examine/pkg/repositories/examine_repository"
	"github.com/kaiiorg/go-bif-examine/pkg/storage"
	"net"

	"github.com/kaiiorg/go-bif-examine/pkg/config"
	"github.com/kaiiorg/go-bif-examine/pkg/rpc/pb"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedBifExamineServer

	config     *config.Config
	log        zerolog.Logger
	grpcServer *grpc.Server

	storage           storage.BifStorage
	examineRepository examine_repository.ExamineRepository
}

func New(examineRepository examine_repository.ExamineRepository, storage storage.BifStorage, conf *config.Config, log zerolog.Logger) *Server {
	s := &Server{
		config:            conf,
		log:               log,
		grpcServer:        grpc.NewServer(),
		examineRepository: examineRepository,
		storage:           storage,
	}

	return s
}

func (s *Server) Run() error {
	hostConn, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Grpc.Port))
	if err != nil {
		return err
	}

	go func() {
		pb.RegisterBifExamineServer(s.grpcServer, s)
		if err := s.grpcServer.Serve(hostConn); err != nil {
			s.log.Fatal().Err(err).Msg("gRPC server stopped!")
		}
	}()

	return nil
}

func (s *Server) Close() {
	s.grpcServer.Stop()
}
