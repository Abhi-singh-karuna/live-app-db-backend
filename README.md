# Live App Database

A Go application that provides a REST API for managing data in a Turso database.

## Features

- CRUD operations for data management
- RESTful API endpoints
- Turso database integration
- Docker support

## Prerequisites

- Go 1.22 or later
- Docker (optional)
- Turso database credentials

## Environment Variables

The following environment variables are required:

- `TURSO_DATABASE_URL`: Your Turso database URL
- `TURSO_AUTH_TOKEN`: Your Turso authentication token

## Running the Application

### Local Development

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set environment variables
4. Run the application:
   ```bash
   go run main.go
   ```

### Using Docker

1. Build the Docker image:

   ```bash
   docker build -t live-app-db .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 \
     -e TURSO_DATABASE_URL=your_database_url \
     -e TURSO_AUTH_TOKEN=your_auth_token \
     live-app-db
   ```

## API Endpoints

- `POST /create` - Create a new data entry
- `GET /read/:id` - Read a data entry by ID
- `PUT /update/:id` - Update a data entry by ID
- `DELETE /delete/:id` - Delete a data entry by ID

## License

MIT
