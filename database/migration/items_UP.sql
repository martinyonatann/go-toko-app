-- Create the items table
CREATE TABLE toko.items (
    item_id SERIAL PRIMARY KEY,
    item_name VARCHAR(255),
    description TEXT,
    price NUMERIC(10, 2),
    created_at TIMESTAMPTZ
);