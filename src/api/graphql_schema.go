package api

import (
	"github.com/graphql-go/graphql"
)

var CoordinatesType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Coordinates",
	Fields: graphql.Fields{
		"lat": &graphql.Field{Type: graphql.Float},
		"lon": &graphql.Field{Type: graphql.Float},
	},
})

var TimeRangeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TimeRange",
	Fields: graphql.Fields{
		"start": &graphql.Field{Type: graphql.String},
		"end":   &graphql.Field{Type: graphql.String},
	},
})

var SurgeDataPointType = graphql.NewObject(graphql.ObjectConfig{
	Name:   "SurgeDataPoint",
	Fields: graphql.Fields{
		// Define fields
	},
})

var surgePredictionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SurgePrediction",
	Fields: graphql.Fields{
		"multiplier":        &graphql.Field{Type: graphql.Float},
		"confidence":        &graphql.Field{Type: graphql.Float},
		"validUntil":        &graphql.Field{Type: graphql.DateTime},
		"optimumPickupTime": &graphql.Field{Type: graphql.DateTime},
	},
})

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
