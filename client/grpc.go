package client

import "google.golang.org/grpc"

func grpcConnect(arg *Arg) (*grpc.ClientConn, error) {
	const fname = "grpcConnect"
	applog.debug(fname, "START")

	conn, err := grpc.Dial("localhost"+arg.Grpc2sPort, grpc.WithInsecure())
	if err != nil {
		applog.error(fname, err)
		return nil, err
	}

	applog.debugf(fname, "サーバとのGRPC接続用にポート %s で待ち受け", arg.Grpc2sPort)
	applog.debug(fname, "END")
	return conn, nil
}
