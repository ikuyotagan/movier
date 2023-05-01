package grpc

import (
	"context"

	movier "github.com/ikuyotagan/movier/pb"
	"github.com/ikuyotagan/movier/pkg/sre/tracing"
)

// ToDo:: доделать если все же будет нужно
func (s *MovierService) GetPersons(ctx context.Context, req *movier.GetPersonsRequest) (*movier.GetPersonsResponse, error) {
	ctx, span := tracing.StartSpan(ctx, "grpc", "movierService.GetPersons")
	defer span.End()

	return &movier.GetPersonsResponse{}, nil
}
