

# ![RealWorld Example App](logo.png)

> ### Go backend with Huma API Framework codebase containing real world examples (CRUD, auth, advanced patterns, etc) that is **based on** the [RealWorld](https://github.com/gothinkster/realworld) spec and API. Some implementation details, endpoints, or features may differ from the official specification.

### [Demo](https://demo.realworld.io/)&nbsp;&nbsp;&nbsp;&nbsp;[RealWorld](https://github.com/gothinkster/realworld)

This codebase was created to demonstrate a fully fledged backend application built with **Go** and **Huma API Framework** including CRUD operations, authentication, routing, database interactions, and more.

It follows many of the patterns and ideas from the RealWorld project, but may differ in some implementation details, endpoints, or features. Please refer to the API and code for specifics.

For more information on the original RealWorld spec and how this works with other frontends/backends, head over to the [RealWorld](https://github.com/gothinkster/realworld) repo.

## Features

- RESTful API following the RealWorld.io specification
- JWT-based authentication system
- Clean architecture with domain-driven design
- PostgreSQL database with Go Jet SQL builder
- Comprehensive error handling
- Goose database migrations
- Configuration management using Viper
- Structured logging using Zerolog
- Middleware support for authentication, logging, and recovery

## Technology Stack

- **Language**: Go 1.22+
- **API Framework**: [Huma v2](https://github.com/danielgtaylor/huma)
- **Web Framework**: Echo (via Huma adapter)
- **Database**: PostgreSQL
- **SQL Builder**: [Jet SQL Builder](https://github.com/go-jet/jet)
- **Migration Tool**: [Goose](https://github.com/pressly/goose)
- **Configuration**: [Viper](https://github.com/spf13/viper)
- **Authentication**: JWT (using [jwx](https://github.com/lestrrat-go/jwx))
- **Logging**: [Zerolog](https://github.com/rs/zerolog)

## Project Structure

```
├── cmd/
│   └── api/          # Application entrypoint
├── config/           # Configuration files and logic
├── gen/              # Generated SQL code (by Jet)
├── internal/         # Internal packages
│   ├── app/          # App initialization and routing
│   ├── ctxkit/       # Context utilities
│   ├── db/           # Database connection and management
│   ├── domain/       # Domain-specific code
│   │   ├── article/  # Article domain
│   │   ├── profile/  # Profile domain
│   │   └── user/     # User domain
│   ├── middlewares/  # HTTP middlewares
│   └── utils/        # Utility functions
├── migrations/       # Database migrations
└── pkg/              # Public packages
    ├── crypt/        # Cryptography utilities
    ├── errs/         # Error handling
    ├── jwtkit/       # JWT utilities
    └── logger/       # Logging utilities
```

## How it Works

The application follows a clean architecture approach with the following layers:

1. **Delivery Layer**: Handles HTTP requests and responses (in `delivery/http` packages)
2. **Use Case Layer**: Contains business logic (in `usecase` packages)
3. **Repository Layer**: Manages data access and persistence (in `repository` packages)
4. **Entity Layer**: Defines domain models (in `entities` packages)

The flow of control typically follows:
- Request → HTTP Handler → Use Case → Repository → Database
- Database → Repository → Use Case → HTTP Handler → Response

## Getting Started

### Option 1: Using DevContainer (Recommended)

This repository comes with a DevContainer configuration that sets up everything you need to get started quickly, including Go 1.22+, PostgreSQL, and all required tools.

#### Prerequisites

- [VS Code](https://code.visualstudio.com/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [DevContainer Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

#### Setup Steps

1. Clone the repository
   ```bash
   git clone https://github.com/gthomas08/realworld-huma.git
   cd realworld-huma
   ```

2. Open the project in VS Code

3. When prompted, click on "Reopen in Container" or use the command palette (F1) and select "Dev Containers: Reopen in Container"

4. VS Code will build and start the DevContainer, which includes:
   - Go 1.22+ development environment
   - PostgreSQL database (preconfigured with admin/root credentials)
   - Goose CLI (for migrations)
   - All necessary VS Code extensions

5. Once the container is running, dependencies will automatically be installed via the `go mod tidy` post-create command

### Option 2: Manual Setup

#### Prerequisites

- Go 1.22 or higher
- PostgreSQL
- Goose CLI (for migrations)

#### Environment Setup

1. Clone the repository
   ```bash
   git clone https://github.com/gthomas08/realworld-huma.git
   cd realworld-huma
   ```

2. Install dependencies
   ```bash
   go mod download
   ```

3. Install Goose (if not already installed)
   ```bash
   go install github.com/pressly/goose/v3/cmd/goose@latest
   ```

4. Set up PostgreSQL database
   ```bash
   # Example using Docker
   docker run --name realworld-postgres -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=root -e POSTGRES_DB=postgres -p 5432:5432 -d postgres
   ```

### Required Steps After Setup

After setting up with either DevContainer or manual setup, you'll need to perform the following steps to get the application fully running:

#### Running Migrations

Run database migrations using the Makefile command:

```bash
# Inside the DevContainer
make migrate DB_HOST=db
```

#### Running the Application

Start the application using the Makefile command:

```bash
make run
```

By default, the application will run on port 8000. Access the API at http://localhost:8000/api

## Development

### Database Code Generation

After making changes to the database schema, regenerate the SQL models:

```bash
# Inside DevContainer
make jet DB_HOST=db

# With manual setup
make jet DB_HOST=localhost
```

## API Documentation

API documentation is available at the `/api/docs` endpoint when the server is running.

## Acknowledgments

- [RealWorld](https://github.com/gothinkster/realworld) for the API spec
- [Huma](https://github.com/danielgtaylor/huma) for the API framework
