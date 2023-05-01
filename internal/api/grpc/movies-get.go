package grpc

import (
	"context"

	movier "github.com/ikuyotagan/movier/pb"
	"github.com/ikuyotagan/movier/pkg/sre/tracing"
)

func (s *MovierService) GetMovies(ctx context.Context, req *movier.GetMoviesRequest) (*movier.GetMoviesResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "grpc", "movierService.SayHello")
	defer span.End()

	movies, err := s.movieManager.GetMovies(ctx, req)
	if err != nil {
		return nil, err
	}

	return &movier.GetMoviesResponse{
		Movies: movies,
	}, nil
}
