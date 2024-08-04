# Go Auth

This is a simple server and Api that registers users and generates a JwT token for authentication.

## Endpoints
- `/`

response: {message: welcome to the go auth app}

- `/ping`

response: {message: server is up and running}

-`/v1/auth/register`
response: {message: user successfully created}
-`/v1/auth/login`
response: {"message": "Login successful", "token": token}

## installation

clone the repo and run

`go run cmd/main.go`
or install air and run
`air`
or for complete functionality
install make and air and run
`make run`