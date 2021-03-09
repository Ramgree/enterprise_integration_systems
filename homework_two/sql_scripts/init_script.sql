CREATE ROLE pgadmin WITH LOGIN PASSWORD 'pgadmin' CREATEDB;
CREATE DATABASE prod;
CREATE DATABASE test;
GRANT ALL PRIVILEGES ON DATABASE prod TO pgadmin;
GRANT ALL PRIVILEGES ON DATABASE test TO pgadmin;
