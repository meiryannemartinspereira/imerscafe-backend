# ☕ imerscafe Backend

Backend of the **imerscafe** project, developed in **Go**.

The service centralizes the deterministic game logic and is responsible for round processing, action validation, and management of the application's official state.

## Responsibilities

The backend is responsible for:

* Executing the formal game rules;
* Validating hard skills;
* Managing the official round state;
* Validating incoming data;
* Comparing selected ingredients with the expected recipe;
* Calculating the score;
* Determining the official round result;
* Consolidating processed data;
* Orchestrating integrations required for round processing.

## Stack

* **Go 1.22+** — backend programming language;
* **net/http** — HTTP server and API layer;
* **encoding/json** — JSON request and response serialization;
* **Go Modules** — dependency management;
* **Go testing** — automated testing using Go's standard testing package.

## Structure

```text
backend/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── handler/
│   │
│   ├── domain/
│   │
│   ├── service/
│   │
│   ├── repository/
│   │
│   └── config/
│
├── tests/
│
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

### `cmd/server/`

Application entry point.

The `main.go` file is responsible for creating and configuring the application, loading configuration, initializing dependencies, and registering the API routes.

### `internal/handler/`

Contains the HTTP handlers used for communication with the client.

Handlers are responsible for:

* Receiving HTTP requests;
* Reading request data;
* Validating request structure;
* Calling application services;
* Returning HTTP responses.

### `internal/domain/`

Contains the domain models and deterministic game rules.

This layer represents the core concepts of the game, such as:

* Rounds;
* Recipes;
* Ingredients;
* Player actions;
* Hard skills;
* Scores;
* Results.

The domain layer contains the rules that determine the official game result.

### `internal/service/`

Contains the application's business logic and use cases.

Responsible for:

* Processing rounds;
* Validating player actions;
* Applying game rules;
* Calculating scores;
* Managing the official round state;
* Returning the processed result.

### `internal/repository/`

Contains abstractions and implementations related to data persistence and external data sources.

Repositories isolate the application and domain logic from infrastructure-specific implementations.

### `internal/config/`

Contains application configuration and environment-related settings.

Configuration is kept separate from the business logic and infrastructure implementations.

## Processing Flow

Round processing takes place entirely in the backend:

```text
Request
   │
   ▼
HTTP Handler
   │
   ▼
Request Validation
   │
   ▼
Business Rules
   │
   ├── Recipe validation
   ├── Ingredient validation
   ├── Hard skill validation
   ├── Score calculation
   └── Round state update
   │
   ▼
Official Result
   │
   ▼
Response
```

The round result is determined exclusively by the backend.

The client does not determine the official score or game result.

## Execution

### Requirements

Make sure Go is installed and available in your environment.

Check the installed version with:

```bash
go version
```

### Install Dependencies

Download the project dependencies:

```bash
go mod download
```

### Run the Application

```bash
go run ./cmd/server
```

The API will be available at:

```text
http://localhost:8080
```

## Build

To build the application:

```bash
go build -o bin/server ./cmd/server
```

Run the generated binary:

```bash
./bin/server
```

## API Documentation

The API documentation depends on the documentation solution configured by the project.

If an interactive API documentation tool is configured, it should be available through the corresponding application endpoint.

## Tests

Run the automated tests with:

```bash
go test ./...
```

To run tests with additional output:

```bash
go test -v ./...
```

To check test coverage:

```bash
go test ./... -cover
```

## Backend Principle

The backend is the source of truth for all deterministic game rules.

```text
Client
   │
   │ Player action data
   ▼
HTTP API
   │
   │ Validation
   ▼
Business Rules
   │
   │ Deterministic result
   ▼
Official Round State
```

**The client provides the action data. The backend validates, processes, and determines the official result.**
