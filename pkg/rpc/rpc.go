package rpc

import (
	"net"
	"fmt"

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
}

func New(conf *config.Config, log zerolog.Logger) *Server {
	s := &Server{
		config:     conf,
		log:        log,
		grpcServer: grpc.NewServer(),
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