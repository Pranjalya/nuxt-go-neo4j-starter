package repository

import (
	"testing"
	"github.com/Pranjalya/nuxt-go-neo4j-starter/server/app/models"
	"github.com/stretchr/testify/assert"
)

func TestParseCypherQueryResult(t *testing.T) { 
	record := RecordMock{}

	movie := models.Movie{}
	ParseCypherQueryResult(record, "m", &movie)

	assert.Equal(t, "Movie title", movie.Title)
}

type RecordMock struct{}

func (r RecordMock) Keys() []string {
	return []string{}
}

func (r RecordMock) Values() []interface{} {
	return nil
}

func (r RecordMock) Get(key string) (interface{}, bool) {
	switch key {
	case "m.title": 
		return "Movie title", true
	default:
		 return "", false
	}
}

func (r RecordMock) GetByIndex(index int) interface{} {
	return nil
}


