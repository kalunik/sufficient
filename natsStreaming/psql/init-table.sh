#!/bin/sh

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
CREATE TABLE IF NOT EXISTS orders (
  uid TEXT PRIMARY KEY,
  data JSONB NOT NULL DEFAULT '{}'::jsonb
);
EOSQL

#	CREATE USER docker;
#	CREATE DATABASE docker;
#	GRANT ALL PRIVILEGES ON DATABASE docker TO docker;

#SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower('dbname');
#
#if [];
#then
#	echo "Database already exists."
##    openrc -s mariadb stop
#else
#	//create db if not exists
#
#	//create user with pass
#
#	//create table uid | jsonb
#fi