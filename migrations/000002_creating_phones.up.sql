CREATE TABLE IF NOT EXISTS phones (
    phone_id SERIAL PRIMARY KEY,
    parent_id INTEGER REFERENCES parents(id) ON DELETE CASCADE,
    phone_number VARCHAR(20) NOT NULL
);
