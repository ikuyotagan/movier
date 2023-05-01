package grpc

import (
	"context"

	movier "github.com/ikuyotagan/movier/pb"
	"github.com/ikuyotagan/movier/pkg/sre/tracing"
)

func (s *MovierService) GetGenres(ctx context.Context, req *movier.GetGenresRequest) (*movier.GetGenresResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "grpc", "movierService.GetGenres")
	defer span.End()

	genres, err := s.movieManager.GetGenres(ctx)
	if err != nil {
		return nil, err
	}

	return &movier.GetGenresResponse{
		Genres: genres,
	}, nil
}
