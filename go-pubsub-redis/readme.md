# Gin-Redis-WebSocket Pub/Sub Boilerplate

## Overview
This is a Go boilerplate built with the **Gin** framework, integrating **Redis** for Pub/Sub and WebSocket connections. It provides a foundation to build real-time applications that involve event-based communication between a WebSocket server and a Redis-backed message broker.

## Features
- **Gin Framework** for handling HTTP routes and WebSocket upgrades.
- **WebSocket Server** for managing real-time connections.
- **Redis Integration** to handle Pub/Sub messaging for broadcasting events to subscribed WebSocket clients.
- **Modular structure** with clear separation of concerns for easier maintainability and scalability.

## Project Structure
```
/project-root
│   main.go
│   go.mod
│   go.sum
│
├───handlers
│       publish.go       # HTTP handler to publish messages to Redis
│       user_handler.go  # Basic HTTP route handlers
│
├───libraries
│   └───redis
│       redis.go         # Redis client setup
│
├───routes
│       routes.go        # Gin HTTP route definitions
│
└───websocket
    ├── utils.go         # WebSocket helper functions
    └── websocket.go     # WebSocket handler and connection management
```

## Getting Started

### Prerequisites
- **Go** (v1.18+ recommended)
- **Redis** (running locally or on a server)
- **Git**

### Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd project-root
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run Redis**:
   - Make sure Redis is running locally or accessible at the specified address in the configuration (by default: `localhost:6379`).

4. **Run the server:**
   ```bash
   go run main.go
   ```

5. **The server will start on `localhost:8080` by default**.

### Environment Variables

If you need to configure Redis or other variables, set them up as environment variables or directly modify the Redis client configuration in the `redis.go` file.

### Endpoints

#### 1. **WebSocket Connection**
- **Endpoint**: `/ws`
- **Method**: `GET`
- **Description**: Establishes a WebSocket connection that listens for incoming messages and manages subscriptions.
- **Message Format**:
  ```json
  {
      "action": "subscribe",
      "event": "1234"
  }
  ```

#### 2. **Publish Messages to Redis**
- **Endpoint**: `/publish`
- **Method**: `GET`
- **Query Params**:
  - `event`: Event name or channel (e.g., `1234`)
  - `message`: Message content to be published
- **Example Request**:
  ```bash
  curl "http://localhost:8080/publish?event=1234&message=Hello+Redis"
  ```
- **Response**:
  ```json
  {
      "status": "Message published successfully",
      "event": "1234",
      "message": "Hello Redis"
  }
  ```

### Code Overview

#### 1. **WebSocket Integration**
The WebSocket server is integrated with Gin and manages subscriptions to events using a Pub/Sub model. Users can subscribe/unsubscribe to specific events, and messages are sent in real-time.

#### 2. **Redis Pub/Sub**
The project uses Redis as a message broker. Events are published to specific Redis channels and forwarded to all subscribed WebSocket clients.

#### 3. **Gin Framework**
Gin is used for handling HTTP routes, WebSocket upgrades, and API endpoints, making it easy to extend the application with more features.

### Contributing
Feel free to open issues, fork the repository, and submit pull requests for improvements and bug fixes.

