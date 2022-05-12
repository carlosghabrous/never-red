-- stop executing if one query could not run successfully
\set ON_ERROR_STOP on

-- get env vars
\set NEVER_RED_DB_NAME `echo $NEVER_RED_DB_NAME`
\set NEVER_RED_ADMIN_USER `echo $NEVER_RED_ADMIN_USER`
\set NEVER_RED_USER `echo $NEVER_RED_USER`
\set NEVER_RED_PWD `echo $NEVER_RED_PWD`

-- connect to neverred db with never_red_admin user
\connect :NEVER_RED_DB_NAME :NEVER_RED_ADMIN_USER

-- create never_red user
CREATE ROLE :NEVER_RED_USER SUPERUSER LOGIN PASSWORD :'NEVER_RED_PWD';
