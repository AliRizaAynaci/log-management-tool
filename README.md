# Log Management Tool

## Project Description

This project is a simple log management tool that collects and stores application logs while providing query capabilities. The developer will build a system to store log data in memory and implement functionality to query this data based on time intervals or levels (INFO, WARN, ERROR). This tool helps in understanding how logging systems work and how log data is stored and managed.

## Core Features

- **Log Storage:** Store application logs in memory.
- **Log Addition:** Allow users to add logs to the system.
- **Log Querying:** Query logs based on time intervals or severity levels.
- **Testability:** Unit tests to verify API functionality.

## Technical Requirements

- Go 1.20 or newer version

## API Endpoints

### Adding Logs
- **Endpoint:** `POST /logs`
- **Request Body:**
  ```json
  {
    "level": "INFO",
    "message": "User logged in",
    "source": "auth-service"
  }
  ```

### Viewing Logs
- **Endpoint:** `GET /logs`
- **Response Body:**
  ```json
  [
    {
      "timestamp": "2024-11-24T12:34:56Z",
      "level": "INFO",
      "message": "User logged in",
      "source": "auth-service"
    }
  ]
  ```

## Planned Features

* **Database Integration:** Add integration with SQL or NoSQL databases (such as PostgreSQL or MongoDB) instead of in-memory storage.
* **Real-time Log Monitoring:** Implement real-time log monitoring capability using WebSocket.
* **Docker Support:** Make the application deployable in a Docker container.