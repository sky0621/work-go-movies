package server

import (
	"log"
	"net"

	moviesc2s "github.com/sky0621/work-go-movies/grpcc2s"
	moviess2p "github.com/sky0621/work-go-movies/grpcs2p"
	"google.golang.org/grpc"
)

// 対 クライアント
func grpcListen(arg *Arg, s2pClient moviess2p.MovieS2PServiceClient) error {
	const fname = "grpcListen"
	applog.debug(fname, "START")

	lis, err := net.Listen("tcp", arg.Grpc2cPort)
	if err != nil {
		applog.error(fname, err)
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	moviesc2s.RegisterMovieC2SServiceServer(grpcServer, MovieHandler{S2pClient: s2pClient})
	grpcServer.Serve(lis)

	applog.debugf(fname, "クライアントとのGRPC接続用にポート %s で待ち受け", arg.Grpc2cPort)
	applog.debug(fname, "END")
	return nil
}

// 対 永続化層
func grpcConnect(arg *Arg) (*grpc.ClientConn, error) {
	const fname = "grpcConnect"
	applog.debug(fname, "START")

	conn, err := grpc.Dial("localhost"+arg.Grpc2pPort, grpc.WithInsecure())
	if err != nil {
		applog.error(fname, err)
		return nil, err
	}

	applog.debugf(fname, "永続化層とのGRPC接続用にポート %s で待ち受け", arg.Grpc2pPort)
	applog.debug(fname, "END")
	return conn, nil
}
