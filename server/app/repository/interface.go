package repository

import (
	"context"

	"github.com/Pranjalya/nuxt-go-neo4j-starter/server/app/graph/model"
	"github.com/Pranjalya/nuxt-go-neo4j-starter/server/app/models"
)

// Repository definition for repository
type Repository interface {
	// Movie
	FindMovieByUUID(ctx context.Context, uuid string) (*models.Movie, error)
	FindMovies(ctx context.Context, title *string, actor *string) ([]*models.Movie, error)
	FindMovieParticipationsByPersonUUID(ctx context.Context, uuid string) ([]*model.Participation, error)
	// Person
	FindPersonByMovieUUID(ctx context.Context, role string, uuid string) ([]*models.Person, error)
}
