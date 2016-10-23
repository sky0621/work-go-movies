package server

import (
	"log"

	moviesc2s "github.com/sky0621/work-go-movies/grpcc2s"
	context "golang.org/x/net/context"
)

// MovieHandler ...
type MovieHandler struct{}

// GetMovie ...
func (h MovieHandler) GetMovie(ctx context.Context, req *moviesc2s.MovieSkey) (*moviesc2s.Movie, error) {

	log.Println("GetPerson!")
	log.Println(req.Skey)
	return sample1(req.Skey), nil
}

// GetMovies ...
func (h MovieHandler) GetMovies(ctx context.Context, req *moviesc2s.Movie) (*moviesc2s.Movies, error) {

	log.Println("GetPersons!")
	log.Println(req.Skey)
	return &moviesc2s.Movies{
		Movies: []*moviesc2s.Movie{sample1(req.Skey), sample2("92011538")},
	}, nil
}

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
