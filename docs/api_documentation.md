# API Documentation

## Overview

This document provides details about the GraphQL API for the predictive surge pricing system.

## GraphQL Schema

### Queries

#### `surgeForecast`

- **Description**: Predicts the surge multiplier for a given route and time window.
- **Arguments**:
  - `origin`: Coordinates (required)
  - `destination`: Coordinates (required)
  - `timeWindow`: TimeRange (required)
- **Returns**: `SurgePrediction`

#### `historicalSurge`

- **Description**: Retrieves historical surge data for a specific location and time range.
- **Arguments**:
  - `location`: Coordinates (required)
  - `timeRange`: TimeRange (required)
- **Returns**: List of `SurgeDataPoint`

### Mutations

#### `lockSurgePrice`

- **Description**: Locks the surge price for a specific route and pickup time.
- **Arguments**:
  - `routeId`: ID (required)
  - `pickupTime`: DateTime (required)
- **Returns**: `PriceLock`

### Subscriptions

#### `surgeUpdate`

- **Description**: Subscribes to updates on surge multipliers for a specific route.
- **Arguments**:
  - `routeId`: ID (required)
- **Returns**: `SurgeUpdate` 