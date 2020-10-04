//go:generate gorunpkg github.com/99designs/gqlgen --verbose

package graph

import (
	"github.com/Pranjalya/nuxt-go-neo4j-starter/server/app/service"
)

// Resolver is the main gql resolver
type Resolver struct {
	Service service.Service
}
