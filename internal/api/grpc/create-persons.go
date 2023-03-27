package grpc

import (
	"context"

	movier "github.com/ikuyotagan/movier/pb"
	"github.com/ikuyotagan/movier/pkg/sre/tracing"
)

func (s *MovierService) CreatePersonsBatch(ctx context.Context, req *movier.CreatePersonsRequest) (*movier.CreatePersonsResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "grpc", "exampleService.SayHello")
	defer span.End()

	exampleInternalBusinessLogicCall(ctx)

	return &movier.CreatePersonsResponse{}, nil
}
