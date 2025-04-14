-- CREATE DATABASE IF NOT EXISTS portfolio;

-- CREATE TABLE IF NOT EXISTS contact_messages (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(255) NOT NULL,
--     email VARCHAR(255) NOT NULL,
--     message TEXT NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

--
-- Check if database exists, and create if it doesn't
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'portfolio') THEN
        PERFORM dblink_exec('dbname=postgres', 'CREATE DATABASE portfolio');
    END IF;
END
$$;

-- Now create the table if it does not exist
CREATE TABLE IF NOT EXISTS contact_messages (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
