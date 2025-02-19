package api

import (
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/pauly7610/uber_surge_prediction/src/models"
	"github.com/stretchr/testify/assert"
)

func TestResolveSurgeForecast(t *testing.T) {
	// Mock input arguments
	args := map[string]interface{}{
		"origin":      models.Coordinates{Lat: 37.7749, Lon: -122.4194},
		"destination": models.Coordinates{Lat: 34.0522, Lon: -118.2437},
		"timeWindow":  models.TimeRange{Start: "2023-01-01T00:00:00Z", End: "2023-01-01T01:00:00Z"},
	}

	// Call resolver
	result, err := resolveSurgeForecast(graphql.ResolveParams{Args: args})

	// Assert no error and valid result
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
