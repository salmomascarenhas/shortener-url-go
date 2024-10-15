# URL Shortener API

This is a simple URL shortener API written in Go. It allows you to shorten URLs and retrieve the original URLs using a generated code. The project uses the Go `net/http` package for creating a web server, and `go-chi/chi` as a router.

## Features
- **POST `/shorten`**: Accepts a URL and returns a shortened code.
- **GET `/{code}`**: Redirects the user to the original URL associated with the given code.

## Requirements
- Go 1.18+ installed on your machine.
- Access to the internet for dependencies installation.

## Getting Started

### Clone the repository

```bash
git clone https://github.com/salmomascarenhas/shortener-url-go.git
cd shortener-url-go
```

### Install dependencies

This project uses a few third-party libraries, namely `go-chi/chi` and `go-chi/middleware`, as well as `math/rand/v2` and `log/slog`. You can install them by running:

```bash
go mod tidy
```

This will automatically download and install all dependencies.

### Run the server

To start the API server, run the following command in the root directory of the project:

```bash
go run main.go
```

The server will be up and running on port `3000`.

### Example requests

- **Shorten a URL**

  You can shorten a URL by sending a POST request to `/shorten` with the following payload:

  ```json
  {
    "url": "https://example.com"
  }
  ```

  Example using `curl`:

  ```bash
  curl -X POST http://localhost:3000/shorten -H "Content-Type: application/json" -d '{"url": "https://example.com"}'
  ```

  Response:
  ```json
  {
    "data": "abc12345"
  }
  ```

- **Retrieve the original URL**

  After shortening a URL, you can retrieve the original by sending a GET request with the shortened code, e.g., `/abc12345`:

  ```bash
  curl http://localhost:3000/abc12345
  ```

  The API will redirect you to the original URL.

## API Endpoints

### POST `/shorten`

- **Description**: Shortens a URL.
- **Request Body**:
  - `url` (string): The original URL that needs to be shortened.
- **Response**:
  - `data` (string): The generated code for the shortened URL.
  - `error` (string): Error message (if applicable).

- **Status Codes**:
  - `201 Created`: URL successfully shortened.
  - `400 Bad Request`: Invalid request body or URL format.

### GET `/{code}`

- **Description**: Redirects to the original URL.
- **Path Parameters**:
  - `code` (string): The shortened URL code.
- **Response**:
  - **Redirects**: To the original URL if found.
  - `404 Not Found`: If the code does not exist in the database.

## Configuration

- The server is configured to run on port `3000` and has the following timeouts:
  - Read Timeout: 10 seconds
  - Write Timeout: 10 seconds
  - Idle Timeout: 1 minute

You can modify these settings in the `run` function located in `main.go`.

## Logging

The API uses Go's `slog` for structured logging. Errors and important events will be logged to the console.

## License

This project is licensed under the MIT License.

---

Feel free to contribute or report issues if you encounter any problems!
