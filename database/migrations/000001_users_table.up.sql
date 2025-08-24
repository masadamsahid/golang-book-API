CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR UNIQUE NOT NULL,
  password VARCHAR NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  created_by VARCHAR,
  modified_at TIMESTAMP DEFAULT NOW(),
  modified_by VARCHAR
);