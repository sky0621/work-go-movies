package main

import (
	"log"
	"net"

	moviesc2s "github.com/sky0621/work-go-movies/grpcc2s"
	context "golang.org/x/net/context"

	"google.golang.org/grpc"
)

func sample1(skey string) *moviesc2s.Movie {
	return &moviesc2s.Movie{
		Skey:          skey,
		Filename:      "MOV0123.mp4",
		Title:         "運動会にて",
		Playtime:      93,
		Photodatetime: 1477160405,
	}
}

func sample2(skey string) *moviesc2s.Movie {
	return &moviesc2s.Movie{
		Skey:          skey,
		Filename:      "MOV0925.mp4",
		Title:         "ハロウィンパーティ",
		Playtime:      114,
		Photodatetime: 1477160607,
	}
}

// GetPersoner ...
type GetPersoner struct{}

// GetPerson ...
func (p GetPersoner) GetPerson(ctx context.Context, req *moviesc2s.MovieSkey) (*moviesc2s.Movie, error) {

	log.Println("GetPerson!")
	log.Println(req.Skey)
	return sample1(req.Skey), nil
}

// GetPersons ...
func (p GetPersoner) GetPersons(ctx context.Context, req *moviesc2s.Movie) (*moviesc2s.Movies, error) {

	log.Println("GetPersons!")
	log.Println(req.Skey)
	return &moviesc2s.Movies{
		Movies: []*moviesc2s.Movie{sample1(req.Skey), sample2("92011538")},
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
	moviesc2s.RegisterMovieC2SServiceServer(grpcServer, GetPersoner{})
	grpcServer.Serve(lis)
}
