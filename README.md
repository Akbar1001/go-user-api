# User Management API

A RESTful API built with GoFiber, PostgreSQL, SQLC, Uber Zap, and Validator.

## Features

* Create User
* Get User by ID
* Update User
* Delete User
* List All Users
* Dynamic Age Calculation
* Request ID Middleware
* Request Duration Logging
* Input Validation
* Structured Logging with Uber Zap

## Tech Stack

* Go
* GoFiber
* PostgreSQL
* SQLC
* Uber Zap
* go-playground/validator

## Project Structure

cmd/server/main.go

config/

db/migrations/

db/sqlc/

internal/

├── handler/

├── repository/

├── service/

├── routes/

├── middleware/

├── models/

└── logger/

## Setup

### Clone Repository

git clone <repository-url>

cd go-user-api

### Install Dependencies

go mod tidy

### Create Database

CREATE DATABASE user_api;

### Configure Environment

Create `.env`

DB_HOST=localhost

DB_PORT=5432

DB_USER=postgres

DB_PASSWORD=your_password

DB_NAME=user_api

SERVER_PORT=5000

### Run Migration

Execute:

CREATE TABLE IF NOT EXISTS users (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL,
dob DATE NOT NULL
);

### Generate SQLC Code

sqlc generate

### Run Application

go run cmd/server/main.go

## API Endpoints

POST /users

GET /users

GET /users/:id

PUT /users/:id

DELETE /users/:id

## Example Response

{
"id": 1,
"name": "Alice",
"dob": "1990-05-10",
"age": 35
}

## Run Tests

go test ./...
