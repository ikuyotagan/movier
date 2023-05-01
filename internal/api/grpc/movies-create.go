package grpc

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	movier "github.com/ikuyotagan/movier/pb"
	"github.com/ikuyotagan/movier/pkg/sre/tracing"
)

func (s *MovierService) CreateMoviesBatch(ctx context.Context, req *movier.CreateMoviesBatchRequest) (*movier.CreateMoviesBatchResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "grpc", "movierService.CreateMoviesBatch")
	defer span.End()

	ids, err := s.movieManager.CreateMovies(ctx, req.Movies)
	if err != nil {
		return nil, err
	}

	return &movier.CreateMoviesBatchResponse{
		Ids: ids,
	}, nil
}

// ExampleInternalBusinessLogicCall is an example of passing ctx and span to internal business logic.
func exampleInternalBusinessLogicCall(ctx context.Context) {
	_, span := tracing.StartCustomSpan(ctx, trace.SpanKindInternal, "service", "someService.Hi")
	defer span.End()
}
