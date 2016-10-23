package persistence

import (
	moviess2p "github.com/sky0621/work-go-movies/grpcs2p"
	context "golang.org/x/net/context"
)

// MovieHandler ...
type MovieHandler struct{}

// GetMovie ...
func (h MovieHandler) GetMovie(ctx context.Context, req *moviess2p.MovieSkey) (*moviess2p.Movie, error) {
	const fname = "GetMovie"
	applog.debug(fname, "START")
	applog.debug(fname, req)
	movie := sample1(req.Skey)
	applog.debug(fname, "END")
	return movie, nil
}

// GetMovies ...
func (h MovieHandler) GetMovies(ctx context.Context, req *moviess2p.Movie) (*moviess2p.Movies, error) {
	const fname = "GetMovies"
	applog.debug(fname, "START")
	applog.debug(fname, req)
	movies := &moviess2p.Movies{
		Movies: []*moviess2p.Movie{sample1(req.Skey), sample2("92011538")},
	}
	applog.debug(fname, "END")
	return movies, nil
}

// TODO 以降はダミー実装。永続化ストレージとの接続ロジック記述後、削除！

func sample1(skey string) *moviess2p.Movie {
	return &moviess2p.Movie{
		Skey:          skey,
		Filename:      "MOV0123b.mp4",
		Title:         "運動会にて2",
		Playtime:      93,
		Photodatetime: 1477160405,
	}
}

func sample2(skey string) *moviess2p.Movie {
	return &moviess2p.Movie{
		Skey:          skey,
		Filename:      "MOV0925b.mp4",
		Title:         "ハロウィンパーティ2",
		Playtime:      114,
		Photodatetime: 1477160607,
	}
}
