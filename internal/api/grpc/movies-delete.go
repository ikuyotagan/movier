package grpc

import (
	"context"

	movier "github.com/ikuyotagan/movier/pb"
	"github.com/ikuyotagan/movier/pkg/sre/tracing"
)

func (s *MovierService) DeleteMovies(ctx context.Context, req *movier.DeleteMoviesRequest) (*movier.DeleteMoviesResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "grpc", "movierService.DeleteMovies")
	defer span.End()
	err := s.movieManager.DeleteMovies(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	return &movier.DeleteMoviesResponse{
		Success: true,
	}, nil
}
