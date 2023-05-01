package grpc

import (
	"context"

	movier "github.com/ikuyotagan/movier/pb"
	"github.com/ikuyotagan/movier/pkg/sre/tracing"
)

func (s *MovierService) CreateGenresBatch(ctx context.Context, req *movier.CreateGenresRequest) (*movier.CreateGenresResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "grpc", "movierService.CreateGenresBatch")
	defer span.End()

	ids, err := s.movieManager.CreateGenres(ctx, req.Genres)
	if err != nil {
		return nil, err
	}

	return &movier.CreateGenresResponse{
		Ids: ids,
	}, nil
}
