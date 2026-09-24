-- Create the schema
CREATE SCHEMA IF NOT EXISTS movies;

-- Set the search path for this session
SET search_path TO movies, PUBLIC;

-- Make the search_path permanent for the database
ALTER DATABASE cinemad SET search_path TO movies, PUBLIC;
