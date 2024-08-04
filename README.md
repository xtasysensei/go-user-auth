
# Go Auth

This is a simple server and API that registers users and generates JWT tokens for authentication.

## Endpoints

- **`/`**
  - **Method:** GET
  - **Response:** `{ "message": "welcome to the go auth app" }`

- **`/ping`**
  - **Method:** GET
  - **Response:** `{ "message": "server is up and running" }`

- **`/v1/auth/register`**
  - **Method:** POST
  - **Request Body:** `{ "username": "your_username","email": "your_email", "password": "your_password", "confirmpassword": "your_password" }`
  - **Response:** `{ "message": "user successfully created" }`

- **`/v1/auth/login`**
  - **Method:** POST
  - **Request Body:** `{ "username": "your_username", "password": "your_password" }`
  - **Response:** `{ "message": "Login successful", "token": "your_jwt_token" }`

## Installation

1. Clone the repository:
    \`\`\`sh
    git clone https://github.com/yourusername/go-auth.git
    cd go-auth
    \`\`\`

2. Run the application:
    \`\`\`sh
    go run cmd/main.go
    \`\`\`

3. Alternatively, use `air` for live reloading (requires installation):
    \`\`\`sh
    air
    \`\`\`

4. For complete functionality, install `make` and `air`, then run:
    \`\`\`sh
    make run
    \`\`\`

## Docker

You can also run the application using Docker:

1. Build the Docker image:
    \`\`\`sh
    docker build -t go-auth .
    \`\`\`

2. Run the Docker container:
    \`\`\`sh
    docker run -p 8000:8000 --env-file .env go-auth
    \`\`\`

## Docker Compose

To run the application along with a PostgreSQL database using Docker Compose:

1. Ensure Docker Compose is installed.
2. Run the following command:
    \`\`\`sh
    docker-compose up --build
    \`\`\`

## Dependencies

- [Chi](https://github.com/go-chi/chi) for routing
- [JWT](https://github.com/dgrijalva/jwt-go) for token generation
- [PostgreSQL](https://www.postgresql.org/) for database
- [Air](https://github.com/air-verse/air): Optional (can be installed with `go install github.com/air-verse/air@latest`)

## Environment Variables

- `SERVER_PORT`: Port on which the server runs (default: 8000)
- `POSTGRES_SERVER`: PostgreSQL server address (default: localhost)
- `POSTGRES_PORT`: PostgreSQL server port (default: 5432)
- `POSTGRES_DB`: PostgreSQL database name (default: go_app)
- `POSTGRES_USER`: PostgreSQL database user (default: sensei)
- `POSTGRES_PASSWORD`: PostgreSQL database password (default: 12345)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
