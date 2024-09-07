#!/bin/bash
set -e
export PGPASSWORD=password;
psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "postgres" <<-EOSQL
  CREATE DATABASE stori;
  GRANT ALL PRIVILEGES ON DATABASE stori TO postgres;
EOSQL