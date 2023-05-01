package models

import movier "github.com/ikuyotagan/movier/pb"

type (
	Movie             = movier.Movie
	MovieCreateParams = movier.CreateMovieParams
	MoviesFilter      = movier.GetMoviesRequest

	ExternalInfo = movier.ExternalInfo
	Rating       = movier.MPAARating

	Person        = movier.Person
	PersonsFilter = movier.GetPersonsRequest

	Genre        = movier.Genre
	GenresFilter = movier.GetGenresRequest
)
