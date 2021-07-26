#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE TABLE events.processed_events (
	id uuid NOT NULL,
	processed_timestamp timestamp NOT NULL,
	event_name varchar(256) NOT NULL
);
EOSQL
