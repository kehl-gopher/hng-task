# HNG Stage 0 Backend Task - Public API

## Description
This is a simple public API built using **Go** for the HNG Stage 0 Backend Task. The API returns basic information such as:
- My registered email address on the HNG12 Slack workspace.
- The current datetime in **ISO 8601** format (UTC).
- The GitHub URL of this project's codebase.

The API is publicly accessible and follows RESTful conventions, handling **CORS** appropriately.

## Features
- Accepts **GET** requests.
- Returns JSON responses in the correct format.
- The `current_datetime` field is dynamically generated in **ISO 8601** format (UTC).
- Deployed to a publicly accessible endpoint.
- Fast response time (< 500ms).

## API Documentation

### **Base URL**
```
https://widespread-odelinda-hng-cc23a786.koyeb.app
```

### **Endpoints**

#### **GET /**
Returns a welcome message.

**Response:**
```json
{
  "message": "Please visit the /me endpoint for user data"
}
```

#### **GET /me**
Returns user details, including email, current datetime (UTC), and GitHub repository URL.

**Response:**
```json
{
  "email": "kelanidarasimi9@gmail.com",
  "current_datetime": "2025-01-30T09:30:00Z",
  "github_url": "https://github.com/kehl-gopher/hng-task"
}
```

## Installation & Setup

### **Prerequisites**
- Go installed ([Download Go](https://go.dev/dl/))
- Git installed ([Download Git](https://git-scm.com/downloads))

### **Running Locally**
1. Clone the repository:
   ```sh
   git clone https://github.com/kehl-gopher/hng-task.git
   ```
2. Navigate to the project folder:
   ```sh
   cd hng-task
   ```
3. Install dependencies (if any):
   ```sh
   go mod tidy
   ```
4. Run the server:
   ```sh
   go run main.go
   ```
5. The API will be available at:
   ```sh
   http://localhost:8080/me
   ```

## Deployment
The API is deployed and accessible at:
```
https://widespread-odelinda-hng-cc23a786.koyeb.app/me
```

## Technologies Used
- **Programming Language:** Go
- **Web Framework:** `net/http`
- **Hosting:** Koyeb

## Hiring Backlink
Interested in hiring a **Golang Developer**? Check out:
[Hire Golang Developers](https://hng.tech/hire/golang-developers)

---

**Author:** Kelani Darasimi ([@kehl-gopher](https://github.com/kehl-gopher))
