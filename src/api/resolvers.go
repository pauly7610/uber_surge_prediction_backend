package api

import (
	"context"
	"time"
	"github.com/graphql-go/graphql"
)

func resolveSurgeForecast(p graphql.ResolveParams) (interface{}, error) {
	// Extract arguments
	origin := p.Args["origin"].(Coordinates)
	destination := p.Args["destination"].(Coordinates)
	timeWindow := p.Args["timeWindow"].(TimeRange)

	// Call prediction service
	prediction, err := services.GetSurgePrediction(origin, destination, timeWindow)
	if err != nil {
		return nil, err
	}

	return prediction, nil
}

func resolveHistoricalSurge(p graphql.ResolveParams) (interface{}, error) {
	// Extract arguments
	location := p.Args["location"].(Coordinates)
	timeRange := p.Args["timeRange"].(TimeRange)

	// Query historical data
	dataPoints, err := services.GetHistoricalSurgeData(location, timeRange)
	if err != nil {
		return nil, err
	}

	return dataPoints, nil
} 