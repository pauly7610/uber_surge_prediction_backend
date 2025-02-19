-- Example SQL for partitioning
CREATE TABLE surge_data (
    id SERIAL PRIMARY KEY,
    location GEOGRAPHY(POINT),
    timestamp TIMESTAMPTZ,
    multiplier FLOAT
) PARTITION BY RANGE (timestamp);

CREATE TABLE surge_data_2023 PARTITION OF surge_data
FOR VALUES FROM ('2023-01-01') TO ('2024-01-01'); 