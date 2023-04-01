-- 1. Once docker compose is up & running. Make sure postgres is also running
--   # docker compose exec db psql -U postgres

-- 2. Create new user for postgres db
CREATE USER vikas_cac WITH PASSWORD 'password';

-- 2. Create Database
CREATE DATABASE platform_cac;

-- 3. Grant privileges to new user
GRANT ALL PRIVILEGES ON DATABASE platform_cac TO vikas_cac;

-- 4. Connect to postgres using PGAdmin client installed as docker container or docker componse from pgadmin GUI (localhost:5050)
-- # docker inspect <container_id> (NetworkSettings.Network section, IPAddress)
