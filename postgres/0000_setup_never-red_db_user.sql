-- Stop execution if error is found
\set ON_ERROR_STOP on

-- get env vars
\set NEVER_RED_DB_NAME `echo $NEVER_RED_DB_NAME`
\set NEVER_RED_ADMIN_USER `echo $NEVER_RED_ADMIN_USER`
\set NEVER_RED_ADMIN_PWD  `echo $NEVER_RED_ADMIN_PWD`

-- create database
\CREATE DATABASE :NEVER_RED_DB_NAME;
\CREATE ROLE :NEVER_RED_ADMIN_USER SUPERUSER LOGIN PASSWORD :'NEVER_RED_ADMIN_PWD';
\CREATE ROLE never_red_read;

ALTER DATABASE neverred SET search_PATH = never_red,public;