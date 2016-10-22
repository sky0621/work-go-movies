package main

import (
	"log"
	"net"

	movies "github.com/sky0621/work-go-movies"
	context "golang.org/x/net/context"

	"google.golang.org/grpc"
)

func sample1(req *movies.ReqMovie) *movies.Movie {
	return &movies.Movie{
		Skey:          req.Skey,
		Filename:      "MOV0123.mp4",
		Title:         "運動会にて",
		Playtime:      93,
		Photodatetime: 1477160405,
	}
}

func sample2(req *movies.ReqMovie) *movies.Movie {
	return &movies.Movie{
		Skey:          "33221144",
		Filename:      "MOV0925.mp4",
		Title:         "ハロウィンパーティ",
		Playtime:      114,
		Photodatetime: 1477160607,
	}
}

// GetPersoner ...
type GetPersoner struct{}

// GetPerson ...
func (p GetPersoner) GetPerson(ctx context.Context, req *movies.ReqMovie) (*movies.Movie, error) {

	log.Println("GetPerson!")
	log.Println(req.Skey)
	return sample1(req), nil
}

// GetPersons ...
func (p GetPersoner) GetPersons(ctx context.Context, req *movies.ReqMovie) (*movies.Movies, error) {

	log.Println("GetPersons!")
	log.Println(req.Skey)
	return &movies.Movies{
		Movies: []*movies.Movie{sample1(req), sample2(req)},
	}, nil
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
