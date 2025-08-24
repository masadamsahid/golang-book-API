CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  created_by VARCHAR NOT NULL,
  modified_at TIMESTAMP WITH TIME ZONE,
  modified_by VARCHAR
);

CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  title VARCHAR NOT NULL,
  description TEXT,
  image_url VARCHAR,
  release_year INTEGER NOT NULL,
  price INTEGER,
  total_page INTEGER NOT NULL,
  thickness VARCHAR NOT NULL,
  category_id INTEGER NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  created_by VARCHAR NOT NULL,
  modified_at TIMESTAMP WITH TIME ZONE,
  modified_by VARCHAR,
  CONSTRAINT fk_books_category FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
);