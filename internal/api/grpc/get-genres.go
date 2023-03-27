package grpc

import (
	"context"

	movier "github.com/ikuyotagan/movier/pb"
	"github.com/ikuyotagan/movier/pkg/sre/tracing"
)

func (s *MovierService) GetGenres(ctx context.Context, req *movier.GetGenresRequest) (*movier.GetGenresResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "grpc", "exampleService.SayHello")
	defer span.End()

	exampleInternalBusinessLogicCall(ctx)

	return &movier.GetGenresResponse{}, nil
}
