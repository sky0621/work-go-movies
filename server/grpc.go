package server

import (
	"log"
	"net"

	moviesc2s "github.com/sky0621/work-go-movies/grpcc2s"
	"google.golang.org/grpc"
)

func grpcListen(arg *Arg) error {
	const fname = "grpcConnect"
	applog.debug(fname, "START")

	lis, err := net.Listen("tcp", arg.GrpcPort)
	if err != nil {
		applog.error(fname, err)
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	moviesc2s.RegisterMovieC2SServiceServer(grpcServer, MovieHandler{})
	grpcServer.Serve(lis)

	applog.debug(fname, "END")
	return nil
}
