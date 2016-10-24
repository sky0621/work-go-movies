package persistence

import (
	moviess2p "github.com/sky0621/work-go-movies/grpcs2p"
	context "golang.org/x/net/context"
)

// MovieHandler ...
type MovieHandler struct {
	storage *storage
}

// GetMovie ...
func (h MovieHandler) GetMovie(ctx context.Context, req *moviess2p.MovieSkey) (*moviess2p.Movie, error) {
	const fname = "GetMovie"
	applog.debug(fname, "START")
	applog.debug(fname, req)

	// var s2pMovies moviess2p.Movie
	// query := h.storage.conn.DB(storageName).C(collectionName).Find(bson.M{})
	// query.One(s2pMovies)
	movie, err := h.storage.readOne(req.Skey)
	if err != nil {
		applog.debug(fname, "ABEND")
		return nil, err
	}
	applog.debug(fname, "END")
	return movie, nil
}

// GetMovies ...
func (h MovieHandler) GetMovies(ctx context.Context, req *moviess2p.Movie) (*moviess2p.Movies, error) {
	const fname = "GetMovies"
	applog.debug(fname, "START")
	applog.debug(fname, req)
	movies, err := h.storage.read()
	if err != nil {
		applog.debug(fname, "ABEND")
		return nil, err
	}
	applog.debug(fname, "END")
	return &moviess2p.Movies{Movies: movies}, nil
}
