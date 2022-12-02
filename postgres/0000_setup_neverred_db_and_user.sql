-- Stop execution if error is found
\set ON_ERROR_STOP on

-- get env vars
\set NEVER_RED_DB_NAME `echo $POSTGRES_DB`
\set NEVER_RED_ADMIN_USER `echo $POSTGRES_USER`
\set NEVER_RED_ADMIN_PWD  `echo $POSTGRES_PASSWORD`

-- create database
-- CREATE DATABASE :NEVER_RED_DB_NAME;
CREATE ROLE :NEVER_RED_ADMIN_USER SUPERUSER LOGIN PASSWORD :'NEVER_RED_ADMIN_PWD';
CREATE ROLE never_red_read;

ALTER DATABASE :NEVER_RED_DB_NAME SET search_path = never_red,public;