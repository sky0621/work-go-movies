package main

import (
	"log"
	"net"

	movies "github.com/sky0621/work-go-movies"
	context "golang.org/x/net/context"

	"google.golang.org/grpc"
)

// GetPersoner ...
type GetPersoner struct{}

// GetPerson ...
func (p GetPersoner) GetPerson(ctx context.Context, req *movies.ReqMovie) (*movies.Movie, error) {

	log.Println("GetPerson!")
	log.Println(req.Id)
	return nil, nil
}

// GetPersons ...
func (p GetPersoner) GetPersons(ctx context.Context, req *movies.ReqMovie) (*movies.Movies, error) {

	log.Println("GetPersons!")
	log.Println(req.Id)
	return nil, nil
}

// server
func main() {
	// clientからの接続のため、ひとまずここに書く
	lis, err := net.Listen("tcp", ":7110")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	movies.RegisterMovieServiceServer(grpcServer, GetPersoner{})
	grpcServer.Serve(lis)
}
