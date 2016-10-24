package server

import (
	moviesc2s "github.com/sky0621/work-go-movies/grpcc2s"
	moviess2p "github.com/sky0621/work-go-movies/grpcs2p"
	context "golang.org/x/net/context"
)

// MovieHandler ...
type MovieHandler struct {
	S2pClient moviess2p.MovieS2PServiceClient // 永続化層との接続情報を保持
}

// GetMovie ... クライアントからの動画リソース要求
func (h MovieHandler) GetMovie(ctx context.Context, req *moviesc2s.MovieSkey) (*moviesc2s.Movie, error) {
	const fname = "GetMovie"
	applog.debug(fname, "START")
	applog.debug(fname, req)
	movie, err := convertedGetMovie(ctx, req, h.S2pClient)
	if err != nil {
		applog.error(fname, err)
		applog.debug(fname, "ABEND")
		return nil, err
	}
	applog.debug(fname, "END")
	return movie, nil
}

// GetMovies ... クライアントからの動画リソース要求
func (h MovieHandler) GetMovies(ctx context.Context, req *moviesc2s.Movie) (*moviesc2s.Movies, error) {
	const fname = "GetMovies"
	applog.debug(fname, "START")
	applog.debug(fname, req)
	movies, err := convertedGetMovies(ctx, req, h.S2pClient)
	if err != nil {
		applog.error(fname, err)
		applog.debug(fname, "ABEND")
		return nil, err
	}
	applog.debug(fname, "END")
	return movies, nil
}

func convertedGetMovie(ctx context.Context, req *moviesc2s.MovieSkey, s2pClient moviess2p.MovieS2PServiceClient) (*moviesc2s.Movie, error) {
	// リクエスト変換
	var reqS2p moviess2p.MovieSkey
	reqS2p.Skey = req.Skey
	// 接続情報を保持していた永続化層のGetMovieを呼び出し
	resS2p, err := s2pClient.GetMovie(ctx, &reqS2p)
	if err != nil {
		return nil, err // TODO エラーログ等
	}
	// レスポンス変換
	var resC2s moviesc2s.Movie
	resC2s.Skey = resS2p.Skey
	resC2s.Filename = resS2p.Filename
	resC2s.Title = resS2p.Title
	resC2s.Playtime = resS2p.Playtime
	resC2s.Photodatetime = resS2p.Photodatetime
	return &resC2s, err
}

func convertedGetMovies(ctx context.Context, req *moviesc2s.Movie, s2pClient moviess2p.MovieS2PServiceClient) (*moviesc2s.Movies, error) {
	const fname = "convertedGetMovies"
	// リクエスト変換
	var reqS2p moviess2p.Movie
	reqS2p.Skey = req.Skey
	reqS2p.Filename = req.Filename
	reqS2p.Title = req.Title
	reqS2p.Playtime = req.Playtime
	reqS2p.Photodatetime = req.Photodatetime
	// 接続情報を保持していた永続化層のGetMovieを呼び出し
	resS2p, err := s2pClient.GetMovies(ctx, &reqS2p)
	if err != nil {
		return nil, err // TODO エラーログ等
	}
	// レスポンス変換
	var resC2sMovies []*moviesc2s.Movie
	for _, resS2pMovie := range resS2p.Movies {
		var resC2s moviesc2s.Movie
		applog.debug(fname, resS2pMovie)
		resC2s.Skey = resS2pMovie.Skey
		resC2s.Filename = resS2pMovie.Filename
		resC2s.Title = resS2pMovie.Title
		resC2s.Playtime = resS2pMovie.Playtime
		resC2s.Photodatetime = resS2pMovie.Photodatetime
		resC2sMovies = append(resC2sMovies, &resC2s)
	}
	return &moviesc2s.Movies{Movies: resC2sMovies}, err
}
