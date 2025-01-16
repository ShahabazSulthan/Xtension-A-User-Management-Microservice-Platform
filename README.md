# Xtension: A User Management Microservice Platform

Xtension is a microservice-based user management platform built with Golang, featuring Docker containerization and Kubernetes deployment. It integrates PostgreSQL and Redis for efficient data management and caching, and supports gRPC communication between services. This project demonstrates scalable architecture, REST API design, and container orchestration.

---

## Features

### Microservice 1: User Management

- **Create a User**: Accepts user details (name, email, phone) and stores them in PostgreSQL.
- **Retrieve User by ID**: Checks Redis for cached data or fetches from PostgreSQL.
- **Update User**: Updates both PostgreSQL and Redis with modified details.
- **Delete User**: Removes user data from PostgreSQL and clears cache in Redis.

### Microservice 2: Parallel and Sequential Methods

- **Method 1**: Processes requests sequentially. Fetches user names from the database and waits for a specified time.
- **Method 2**: Processes requests in parallel, independent of Method 1. Fetches user names and waits for the specified time.
- **gRPC Integration**: Seamless communication between the two microservices.

---

## Directory Structure

```
├── Deployment
│   ├── microservice_one.yaml      # Kubernetes manifest for Microservice 1
│   ├── microservice_two.yaml      # Kubernetes manifest for Microservice 2
│   ├── Postgres.yaml              # PostgreSQL deployment
│   └── redis.yaml                 # Redis deployment
├── Microservice=1
│   ├── cmd/main.go                # Entry point for Microservice 1
│   ├── pkg
│   │   ├── config                 # Configuration management
│   │   ├── db                     # Database and Redis connection setup
│   │   ├── di                     # Dependency injection
│   │   ├── model                  # Data models
│   │   ├── pb                     # Protocol Buffers for gRPC
│   │   ├── repo                   # Repository layer for database interaction
│   │   ├── server                 # gRPC server
│   │   └── usecase                # Business logic implementation
├── Microservice=2
│   ├── cmd/main.go                # Entry point for Microservice 2
│   ├── pkg
│   │   ├── config                 # Configuration management
│   │   ├── Microservice           # Client, handler, model, and routes
│   │   ├── pb                     # Protocol Buffers for gRPC
│   │   ├── Queue                  # Queue handling (e.g., for parallel tasks)
│   │   └── usecase                # Business logic implementation
├── compose.yaml                   # Docker Compose file
└── .gitignore                     # Git ignore file
```

---

## Setup Instructions

### Prerequisites

- Golang (v1.20 or above)
- Docker
- Kubernetes (Minikube or any other cluster)
- PostgreSQL and Redis

### Local Development

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd Xtension
   ```
2. Configure environment variables in `dev.env` for each microservice.
3. Run Microservice 1:
   ```bash
   cd Microservice=1
   go run cmd/main.go
   ```
4. Run Microservice 2:
   ```bash
   cd Microservice=2
   go run cmd/main.go
   ```

---

## Docker Containerization

1. Build Docker images for each microservice:
   ```bash
   docker build -t microservice1 ./Microservice=1
   docker build -t microservice2 ./Microservice=2
   ```
2. Run the services using Docker Compose:
   ```bash
   docker-compose up
   ```

---

## Kubernetes Deployment

1. Apply the PostgreSQL and Redis deployments:
   ```bash
   kubectl apply -f Deployment/Postgres.yaml
   kubectl apply -f Deployment/redis.yaml
   ```
2. Deploy Microservices:
   ```bash
   kubectl apply -f Deployment/microservice_one.yaml
   kubectl apply -f Deployment/microservice_two.yaml
   ```
3. Verify the deployments:
   ```bash
   kubectl get pods
   ```
4. Expose the services using a LoadBalancer or NodePort:
   ```bash
   kubectl expose deployment microservice-one --type=NodePort --port=8080
   kubectl expose deployment microservice-two --type=NodePort --port=8081
   ```

---

## API Documentation

### User Routes (Microservice 1)

- `POST /users/create` - Create a new user
- `GET /users/:id` - Retrieve a user by ID
- `PATCH /users/update` - Update user details
- `DELETE /users/delete/:id` - Delete a user

### Method Routes (Microservice 2)

- `POST /users/method` - Execute Method 1 or Method 2 based on input

#### Sample Request

```json
{
  "method": 1,
  "waitTime": 5
}
```

---

## Additional Notes

- Protocol Buffers (`.proto`) are used for gRPC communication.
- Testing is implemented using `mockgen` and `sqlmock`.
- Follow clean architecture principles.

---

## Contribution Guidelines

- Fork the repository and create a feature branch.
- Ensure code passes all tests before creating a pull request.
- Follow the coding standards and documentation guidelines.

---


## Authors

Shahabaz Sulthan;

