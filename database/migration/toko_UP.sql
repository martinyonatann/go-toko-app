CREATE SCHEMA IF NOT EXISTS toko;

-- Create the users table
CREATE TABLE IF NOT EXISTS toko.users (
    user_id SERIAL PRIMARY KEY,
    fullname VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
-- Create the items table
CREATE TABLE IF NOT EXISTS toko.items (
    item_id SERIAL PRIMARY KEY,
    item_name VARCHAR(255),
    description TEXT,
    price NUMERIC(10, 2),
    created_at TIMESTAMPTZ
);