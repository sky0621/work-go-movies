package persistence

import (
	"log"
	"net"

	moviess2p "github.com/sky0621/work-go-movies/grpcs2p"
	"google.golang.org/grpc"
)

// 対 サーバ
func grpcListen(arg *Arg) error {
	const fname = "grpcListen"
	applog.debug(fname, "START")

	lis, err := net.Listen("tcp", arg.Grpc2sPort)
	if err != nil {
		applog.error(fname, err)
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	moviess2p.RegisterMovieS2PServiceServer(grpcServer, MovieHandler{})
	grpcServer.Serve(lis)

	applog.debugf(fname, "サーバとのGRPC接続用にポート %s で待ち受け", arg.Grpc2sPort)
	applog.debug(fname, "END")
	return nil
}
