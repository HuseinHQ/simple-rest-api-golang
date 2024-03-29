# SIMPLE GOLANG REST API

## Overview

Simple API Documentation for login and register a user with JWT.

## Prerequisites

- Go installed on your machine.
- Any additional dependencies.

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/HuseinHQ/simple-rest-api-go.git
    cd your-api
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Build and run the application:

    ```bash
    go run main.go
    ```

Your API should now be running at `http://localhost:3000` (or another specified port).

## API Endpoints

### 1. User Registration

- **Endpoint**: `POST /register`
- **Request Payload**:
  ```json
  {
    "name": "John Doe",
    "email": "johndoe@mail.com",
    "password": "secret"
  }
  ```
- **Response**:
  - _200_
    ```
    New user registered sucessfully!
    ```
  
  - _400_
    ```
    Invalid email/password
    ```
    OR
    ```
    Failed to create token
    ```

### 2. User Login

- **Endpoint**: `POST /login`
- **Request Payload**:
  ```json
  {
    "email": "example@mail.com",
    "password": "secret"
  }
- **Response**:
  - _200_
    ```json
    "access token"
    ```