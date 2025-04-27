#!/bin/bash
set -e
# this runs inside the Postgres container at startup
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
  CREATE DATABASE blabberit;
EOSQL
