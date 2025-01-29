Here is the updated `README.md` reflecting the exact response format you requested:

---

# HNG12 Backend - Stage 0 Task

Welcome to the backend task for HNG12! This project is a simple API developed using Go, which returns basic information about the developer.

## Project Description

This API returns:
- The registered email address (used to register on the HNG12 Slack workspace).
- The current datetime in ISO 8601 format (UTC).
- The GitHub URL of the project’s codebase.

## Tech Stack

- **Programming Language/Framework**: Go
- **Deployment**: The API is hosted on a publicly accessible endpoint.

## API Documentation

### Endpoint

- **URL**: `GET <your-url>`

### Request

- Method: `GET`
- No request body is required.

### Response (200 OK)

```json
{
  "email": "kelanidarasimi9@gmail.com",
  "current_datetime": "2025-01-29T13:02:28Z",
  "github_url": "https://github.com/kehl-gopher/hng-task"
}
```

### Response Fields:

- **email**: Your registered email address.
- **current_datetime**: The current datetime in ISO 8601 format (UTC).
- **github_url**: The URL of your project’s GitHub repository.

## Deployment

The API is deployed and can be accessed at a publicly available URL. You can access it by visiting the endpoint:

- `https://widespread-odelinda-hng-cc23a786.koyeb.app/me`

## Code Structure

- **main.go**: The main entry point for the application where the server and routes are set up.
- **README.md**: This file.
- **go.mod**: The Go module file for dependencies.

## Example Usage

To get the information, simply send a `GET` request to the endpoint:

```bash
curl -X GET <your-url>
```

The response will be in the following format:

```json
{
  "email": "kelanidarasimi9@gmail.com",
  "current_datetime": "2025-01-29T13:02:28Z",
  "github_url": "https://github.com/kehl-gopher/hng-task"
}
```

## Backlink to HNG Tech

For more information on hiring developers in Go, visit: [HNG Tech - Hire Go Developers](https://hng.tech/hire/golang-developers)
