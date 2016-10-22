package client

import "google.golang.org/grpc"

func grpcConnect(arg *Arg) (*grpc.ClientConn, error) {
	const fname = "grpcConnect"
	applog.debug(fname, "START")

	conn, err := grpc.Dial("localhost"+arg.GrpcPort, grpc.WithInsecure())
	if err != nil {
		applog.error(fname, err)
		return nil, err
	}

	applog.debug(fname, "END")
	return conn, nil
}
