# Pi Trainer

A simple web app to train and track how many digits of π you can memorize

## Features

* User registration & login
* Track how many digits of π you can recall
* Save your progress
* Simple leaderboard (future)

## Tech Stack

* Backend: Go (Gin)
* Database: SQLite / PostgreSQL
* Frontend: [deciding]

## Future Ideas

* Global leaderboard
* Timed challenges
* Difficulty levels
* UI improvements

# API Documentation
## Base URL
All endpoints are prefixed with `/api`.

## Authentication
Some endpoints require authentication using JWT tokens. Include the token in the `Authorization` header as `Bearer <token>`.

## Endpoints

### Health Check
- **Endpoint**: `GET /api/health`
- **Description**: Checks if the API is running and healthy.
- **Authentication**: Not required
- **Response**: 
  ```json
  {
    "status": "ok"
  }
  ```

### Create User
- **Endpoint**: `POST /api/users`
- **Description**: Creates a new user account.
- **Authentication**: Not required
- **Request Body**:
  ```json
  {
    "username": "string (required, 3-25 characters)",
    "password": "string (required)"
  }
  ```
- **Response** (Success):
  ```json
  {
    "success": true
  }
  ```
- **Response** (Error):
  ```json
  {
    "success": false,
    "error": "error message"
  }
  ```

### Login
- **Endpoint**: `POST /api/auth/login`
- **Description**: Authenticates a user and returns a JWT token.
- **Authentication**: Not required
- **Request Body**:
  ```json
  {
    "username": "string (required)",
    "password": "string (required)"
  }
  ```
- **Response** (Success):
  ```json
  {
    "token": "jwt_token_string"
  }
  ```
- **Response** (Error):
  ```json
  {
    "error": "invalid credentials"
  }
  ```

### Get Current User
- **Endpoint**: `GET /api/me`
- **Description**: Retrieves information about the currently authenticated user.
- **Authentication**: Required (Bearer token)
- **Response** (Success):
  ```json
  {
    "id": 1,
    "username": "example_user",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
  ```
- **Response** (Error):
  ```json
  {
    "error": "user not found"
  }
  ```