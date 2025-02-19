package api

import (
	"github.com/graphql-go/graphql"
	"github.com/pauly7610/uber_surge_prediction/src/models"
	"github.com/pauly7610/uber_surge_prediction/src/services"
)

func resolveSurgeForecast(p graphql.ResolveParams) (interface{}, error) {
	origin := p.Args["origin"].(models.Coordinates)
	destination := p.Args["destination"].(models.Coordinates)
	timeWindow := p.Args["timeWindow"].(models.TimeRange)

	prediction, err := services.GetSurgePrediction(origin, destination, timeWindow)
	if err != nil {
		return nil, err
	}

	return prediction, nil
}

func resolveHistoricalSurge(p graphql.ResolveParams) (interface{}, error) {
	location := p.Args["location"].(models.Coordinates)
	timeRange := p.Args["timeRange"].(models.TimeRange)

	dataPoints, err := services.GetHistoricalSurgeData(location, timeRange)
	if err != nil {
		return nil, err
	}

	return dataPoints, nil
}
