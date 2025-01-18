# Go Hello World Server

A simple HTTP server written in Go that responds with "Hello World".

## Prerequisites

- Go 1.23 or higher

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   └── handler/
│       └── handler.go
├── test/
│   └── hello_test.go
├── scripts/
│   └── setup.sh
├── Makefile
├── go.mod
└── README.md
```

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/infysumanta/go-hello-world-server.git
   cd go-hello-world-server
   ```

2. Set up the project:

   ```bash
   make setup
   ```

3. Build the server:

   ```bash
   make build
   ```

4. Run the server:
   ```bash
   make run
   ```

The server will start on port 8080. You can access it at http://localhost:8080

## Available Make Commands

- `make setup`: Initialize and set up the project
- `make build`: Build the server binary
- `make run`: Build and run the server
- `make test`: Run the test suite
- `make clean`: Clean up build artifacts

## Testing

Run the tests using:

```bash
make test
```

## API Endpoints

- `GET /`: Returns "Hello World"

## License

MIT License

## Author

Sumanta Kabiraj
