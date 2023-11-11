# Movie CRUD API

This Go application demonstrates a basic RESTful API for a movie database, implementing CRUD operations (Create, Read, Update, Delete) using the Gorilla mux router.

## Usage

### Prerequisites

- Go installed on your machine
- Gorilla mux library (`github.com/gorilla/mux`) installed

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/KarkiAnmol/CRUD_API_with_Golang.git
   ```

2. Navigate to the project directory:

   ```bash
   cd movie-crud-api
   ```

3. Install dependencies:

   ```bash
   go get
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

The server will start at `http://localhost:8084`.

## API Endpoints

### Get all movies

```http
GET /movies
```

### Get a specific movie by ID

```http
GET /movies/{id}
```

### Create a new movie

```http
POST /movies
```

### Update a movie by ID

```http
PUT /movies/{id}
```

### Delete a movie by ID

```http
DELETE /movies/{id}
```

## Sample Movie JSON

```json
{
  "id": "1",
  "isbn": "1234",
  "title": "movie one",
  "director": {
    "firstname": "Director",
    "lastname": "One"
  }
}
```
