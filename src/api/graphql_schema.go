package api

import (
	"github.com/graphql-go/graphql"
)

var surgePredictionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SurgePrediction",
	Fields: graphql.Fields{
		"multiplier": &graphql.Field{Type: graphql.Float},
		"confidence": &graphql.Field{Type: graphql.Float},
		"validUntil": &graphql.Field{Type: graphql.DateTime},
		"optimumPickupTime": &graphql.Field{Type: graphql.DateTime},
	},
})

// Define other types and schema
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"surgeForecast": &graphql.Field{
			Type: surgePredictionType,
			Args: graphql.FieldConfigArgument{
				"origin": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(CoordinatesType),
				},
				"destination": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(CoordinatesType),
				},
				"timeWindow": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(TimeRangeType),
				},
			},
			Resolve: resolveSurgeForecast,
		},
		"historicalSurge": &graphql.Field{
			Type: graphql.NewList(SurgeDataPointType),
			Args: graphql.FieldConfigArgument{
				"location": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(CoordinatesType),
				},
				"timeRange": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(TimeRangeType),
				},
			},
			Resolve: resolveHistoricalSurge,
		},
	},
}) 