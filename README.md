 # Book Management API

A RESTful API for managing a collection of books, built with Go, Gin, and PostgreSQL. This project provides endpoints for user authentication, and for creating, reading, updating, and deleting books and categories.

## Tech Stack

- **Go**: Backend Language
- **Gin**: HTTP web framework
- **PostgreSQL**: Database
- **Go Playground Validator**: For request validation
- **Bcrypt**: For password hashing
- **Golang-JWT**: For JWT authentication

## Getting Started

### Prerequisites

- Go (version 1.24.5 or higher)
- PostgreSQL
- `migrate` CLI tool

### Installation & Setup

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/masadamsahid/golang-book-API.git
    cd golang-book-API
    ```

2.  **Create a `.env` file** in the root directory by copying the `.env.example` file and fill in your database credentials:
    ```
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=your_db_name
    DB_SSLMODE=disable
    ```

3.  **Run database migrations: (using [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate))**
    ```sh
    migrate -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}" -path database/migrations up
    ```

4.  **Install dependencies and run the server:**
    ```sh
    go mod tidy
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

---

## API Endpoint Documentation

### Users

#### `POST /api/users/register`

Registers a new user.

**Request Body:**
```json
{
    "username": "testuser",
    "password": "password123",
    "confirm_password": "password123"
}
```

**Responses:**
- `201 Created`: User successfully registered.
- `400 Bad Request`: Validation error (e.g., passwords don't match, username not alphanumeric).
- `409 Conflict`: Username already exists.

#### `POST /api/users/login`

Logs in a user and returns a JWT.

**Request Body:**
```json
{
    "username": "testuser",
    "password": "password123"
}
```

**Responses:**
- `200 OK`: Login successful, returns a token.
- `400 Bad Request`: Invalid credentials or validation error.

---

### Categories

*(Authentication required for create, update, and delete category endpoints. Non authenticated users are only authorized to read only)*

#### `POST /api/categories`

Creates a new category.

**Request Body:**
```json
{
    "name": "Fiction"
}
```

#### `GET /api/categories`

Retrieves a list of all categories.

#### `GET /api/categories/:id`

Retrieves a single category by its ID.

#### `PUT /api/categories/:id`

Updates an existing category's name.

**Request Body:**
```json
{
    "name": "Science Fiction"
}
```

#### `DELETE /api/categories/:id`

Deletes a category by its ID.

#### `GET /api/categories/:id/books`
Retrieves all books belonging to a specific category.

---

### Books

*(Authentication required for create, update, and delete books endpoints. Non authenticated users are only authorized to read only)*

#### `POST /api/books`

Creates a new book.

**Request Body:**
```json
{
    "title": "Harry Potter and the Sorcerer's Stone",
    "description": "The first book in the Harry Potter series.",
    "image_url": "http://example.com/harry_potter_1.jpg",
    "release_year": 1997,
    "price": 120000,
    "total_page": 309,
    "category_id": 2
}
```




#### `GET /api/books`

Retrieves a list of all books.

#### `GET /api/books/:id`

Retrieves a single book by its ID.

#### `PUT /api/books/:id`

Updates a book's details. Undefined or null field means the field's value will be nulled/removed.

**Request Body (Example):**
```json
{
    "title": "Harry Potter and the Sorcerer's Stone (updated)",
    "description": "The first book in the Harry Potter series. Additional desc",
    "image_url": "http://new-link.com/harry_potter_1.jpg",
    "release_year": 1997,
    "price": 143000,
    "total_page": 89,
    "category_id": 2
}
```

#### `DELETE /api/books/:id`

Deletes a book by its ID.
