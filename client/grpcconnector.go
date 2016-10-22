package client

import (
	"context"
	"log"
	"time"

	movies "github.com/sky0621/work-go-movies"
	"google.golang.org/grpc"
)

func grpcConnect(arg *Arg) int {
	const fname = "grpcConnect"
	applog.debug(fname, "START")

	conn, err := grpc.Dial("localhost"+arg.GrpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		time.Sleep(10 * time.Second)
		conn.Close()
	}()

	client := movies.NewMovieServiceClient(conn)
	r, err := client.GetPerson(context.Background(), &movies.ReqMovie{Id: 12565})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %d", r.Id)

	applog.debug(fname, "END")
	return ExitCodeOK
}
