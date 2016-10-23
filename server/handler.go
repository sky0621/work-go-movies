package server

import (
	moviesc2s "github.com/sky0621/work-go-movies/grpcc2s"
	moviess2p "github.com/sky0621/work-go-movies/grpcs2p"
	context "golang.org/x/net/context"
)

// MovieHandler ...
type MovieHandler struct {
	S2pClient moviess2p.MovieS2PServiceClient
}

// GetMovie ...
func (h MovieHandler) GetMovie(ctx context.Context, req *moviesc2s.MovieSkey) (*moviesc2s.Movie, error) {
	const fname = "GetMovie"
	applog.debug(fname, "START")
	applog.debug(fname, req)
	movie, err := h.GetMovie(ctx, req)
	if err != nil {
		applog.error(fname, err)
		applog.debug(fname, "ABEND")
		return nil, err
	}
	applog.debug(fname, "END")
	return movie, nil
}

// GetMovies ...
func (h MovieHandler) GetMovies(ctx context.Context, req *moviesc2s.Movie) (*moviesc2s.Movies, error) {
	const fname = "GetMovies"
	applog.debug(fname, "START")
	applog.debug(fname, req)
	movies, err := h.GetMovies(ctx, req)
	if err != nil {
		applog.error(fname, err)
		applog.debug(fname, "ABEND")
		return nil, err
	}
	applog.debug(fname, "END")
	return movies, nil
}
