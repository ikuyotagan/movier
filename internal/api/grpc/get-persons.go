package grpc

import (
	"context"
	movier "github.com/ikuyotagan/movier/pb"
	"github.com/ikuyotagan/movier/pkg/sre/tracing"
)

func (s *MovierService) GetPersons(ctx context.Context, req *movier.GetPersonsRequest) (*movier.GetPersonsResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "grpc", "exampleService.SayHello")
	defer span.End()

	exampleInternalBusinessLogicCall(ctx)

	return &movier.GetPersonsResponse{}, nil
}
