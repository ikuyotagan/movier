package grpc

import (
	"github.com/ikuyotagan/movier/internal/business/movies"
	movier "github.com/ikuyotagan/movier/pb"
)

type MovierService struct {
	movier.UnimplementedMovierServiceServer

	movieManager *movies.Manager
}

func NewMovierService(movieManager *movies.Manager) MovierService {
	return MovierService{
		movieManager: movieManager,
	}
}
