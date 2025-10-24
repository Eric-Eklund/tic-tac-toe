# Tic-Tac-Toe Backend Server

A modern Go HTTP server built with Gin framework that serves the React frontend, provides RESTful API endpoints, and manages real-time WebSocket connections for the tic-tac-toe game.

## 🏗️ Architecture

This Go backend serves multiple purposes:

1. **Static File Server**: Serves the built React application and its assets
2. **REST API Server**: Provides endpoints for game state management
3. **WebSocket Server**: Enables real-time communication for multiplayer features
4. **Persistence Layer**: JSON file-based storage for game matches

## 🔧 Tech Stack

- **Go 1.25** - Backend language
- **Gin Framework** - High-performance HTTP web framework
- **Gorilla WebSocket** - WebSocket implementation
- **Google UUID** - Unique game match identifiers
- **Gin CORS** - Cross-Origin Resource Sharing middleware
- **Gin Static** - Static file serving middleware
- **Alpine Linux** - Production container base image

## 📁 Project Structure

```
back-end/
├── internal/              # Internal packages (clean architecture)
│   ├── config/         # Configuration management
│   │   └── config.go   # Environment variables and CORS setup
│   ├── handlers/       # HTTP request handlers
│   │   ├── handler.go          # Main route setup
│   │   ├── health_handler.go   # Health check endpoint
│   │   ├── match_handler.go    # Game match management
│   │   ├── static_handler.go   # Static file serving
│   │   ├── game_board_handler.go
│   │   └── player_handlers.go
│   ├── models/         # Data models
│   │   ├── match.go        # Match structure and constructor
│   │   ├── game_board.go   # Game board model
│   │   └── player.go       # Player model
│   ├── services/       # Business logic
│   │   └── game_service.go # Match persistence (JSON)
│   └── websocket/      # WebSocket infrastructure
│       └── hub.go          # WebSocket hub, client, and handlers
├── main.go             # Application entry point
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── matches.json        # Local game storage (development)
└── README.md           # This file
```

## 🚀 Getting Started

### Prerequisites

- Go 1.25 or higher
- Port 8080 available

### Development

```bash
# Navigate to backend directory
cd back-end

# Download dependencies
go mod download

# Run the development server
go run main.go

# Server starts at http://localhost:8080
# WebSocket available at ws://localhost:8080/ws
```

### Build Binary

```bash
# Build for current platform
go build -o tic-tac-toe-server main.go

# Build for Linux (production)
CGO_ENABLED=0 GOOS=linux go build -o main .
```

## 🔌 API Endpoints

### REST API Endpoints

| Method | Path | Description | Response |
|--------|------|-------------|----------|
| GET | `/api/health` | Health check endpoint | `{"status":"ok","service":"tic-tac-toe-backend"}` |
| GET | `/api/new-match` | Start a new game match | Match object with ID, game board, and players |
| GET | `/` | Serves React application | HTML page |
| GET | `/*` | Serves static assets | CSS, JS, images, etc. |

### WebSocket Endpoints

| Path | Description |
|------|-------------|
| `/ws` | WebSocket connection for real-time updates |

**WebSocket Message Types:**
- `hello` - Client greeting message (server responds with `welcome`)
- Custom game events - Broadcast to all connected clients

### API Response Examples

**GET /api/new-match**
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "game_board": {
    "board": [
      [null, null, null],
      [null, null, null],
      [null, null, null]
    ]
  },
  "players": [
    {"id": 1, "name": "Player 1", "symbol": "X"},
    {"id": 2, "name": "Player 2", "symbol": "O"}
  ]
}
```

## 📂 Static File Serving

The server uses Gin's static middleware for serving React build files:

```go
// From internal/handlers/static_handler.go
router.Use(static.Serve("/", static.LocalFile("./static", false)))
```

**File Structure in Production:**
```
/app/
├── main           # Go binary
├── static/        # React build output
│   ├── index.html
│   ├── assets/
│   │   ├── index-[hash].css
│   │   └── index-[hash].js
│   ├── bg-pattern-dark.png
│   └── vite.svg
└── data/
    └── matches.json  # Persistent storage
```

## 🔧 Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP server port |
| `HOST` | `localhost` | Server host |
| `DEBUG` | `false` | Enable debug mode (verbose logging) |
| `DATA_DIR` | `.` | Directory for data persistence |

### Server Configuration

Configuration is loaded from environment variables in `internal/config/config.go`:

```go
type Config struct {
    Port  string
    Host  string
    Debug bool
}
```

### CORS Configuration

CORS is configured to allow the Vite development server:

```go
// From internal/config/config.go
AllowOrigins: []string{"http://localhost:5173"},
AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
AllowCredentials: true,
```

## 🛠️ Development

### Architecture Overview

The backend follows a clean architecture pattern:

```
main.go → handlers → services → models
                  ↑
            websocket hub
```

### Adding New API Endpoints

1. **Add route in `internal/handlers/handler.go`:**
```go
func (h *Handler) SetupRoutes() *gin.Engine {
    router := gin.Default()
    
    // Add new route
    router.GET("/api/your-endpoint", h.YourHandler)
    
    return router
}
```

2. **Implement handler function:**
```go
func (h *Handler) YourHandler(c *gin.Context) {
    // Your logic here
    c.JSON(http.StatusOK, gin.H{
        "status": "success",
        "data": yourData,
    })
}
```

3. **Add models in `internal/models/` if needed:**
```go
type YourModel struct {
    ID   string `json:"id"`
    Data string `json:"data"`
}
```

### WebSocket Integration

WebSocket infrastructure is in `internal/websocket/hub.go`:

```go
// The hub maintains active clients and broadcasts messages
type Hub struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
}

// Add WebSocket route in main.go
router.GET("/ws", func(c *gin.Context) {
    websocket.ServeWs(hub, c.Writer, c.Request)
})
```

### Data Persistence

Game matches are persisted to JSON in `internal/services/game_service.go`:

```go
func SaveMatch(g *models.Match) {
    // Loads existing matches
    // Appends or updates the match
    // Writes to matches.json
}

func LoadMatches() (*[]models.Match, error) {
    // Reads from matches.json
    // Returns array of matches
}
```

**Storage Location:**
- Development: `./matches.json`
- Production (Docker): `/app/data/matches.json` (volume mounted)

## 🐳 Docker Integration

The backend is built in a multi-stage Docker build:

**Build Stage:**
```dockerfile
FROM golang:1.25-alpine AS backend-build
WORKDIR /app/backend

# Copy go mod files first for better caching
COPY back-end/go.mod back-end/go.sum ./
RUN go mod download

# Copy source code
COPY back-end/ ./

# Build with cache mount for faster rebuilds
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o main .
```

**Production Stage:**
```dockerfile
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=backend-build /app/backend/main .
COPY --from=frontend-build /app/frontend/dist ./static

EXPOSE 8080
CMD ["./main"]
```

**Docker Compose:**
```yaml
services:
  tic-tac-toe:
    build: .
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      - DEBUG=true
      - DATA_DIR=/app/data
    volumes:
      - ./data:/app/data  # Persistent storage
```

## 🧪 Testing

### Manual Testing

**REST API:**
```bash
# Test health endpoint
curl http://localhost:8080/api/health

# Test new match creation
curl http://localhost:8080/api/new-match

# Test static file serving
curl http://localhost:8080/
```

**WebSocket:**
```javascript
// Browser console test
const ws = new WebSocket('ws://localhost:8080/ws');
ws.onopen = () => ws.send(JSON.stringify({type: 'hello', payload: {}}));
ws.onmessage = (e) => console.log('Received:', e.data);
```

## 📊 Performance Considerations

- **Gin Framework**: High-performance HTTP router with minimal overhead
- **Static File Caching**: Gin static middleware includes caching
- **Binary Size**: ~8-10MB statically compiled with all dependencies
- **Memory Usage**: ~10-30MB at runtime (depends on active connections)
- **Concurrency**: Gin and Go's HTTP server handle concurrent requests efficiently
- **WebSocket**: Gorilla WebSocket library provides production-ready performance
- **JSON Storage**: Simple file-based persistence (suitable for low traffic; consider DB for scale)

## 🔍 Troubleshooting

### Common Issues

1. **Port 8080 in use:**
   ```bash
   lsof -i :8080
   kill -9 <PID>
   ```

2. **Static files not found:**
   - Check `./static` directory exists
   - Verify React build completed successfully

3. **API endpoints not responding:**
   - Check route registration in `main.go`
   - Verify handler function exists

### Debugging

```bash
# Run with verbose logging
go run main.go -v

# Check server logs
docker logs tic-tac-toe-container
```

## 🎯 Future Enhancements

- [x] Environment-based configuration
- [x] Gin framework integration
- [x] WebSocket support for real-time features
- [x] CORS configuration
- [x] JSON file persistence
- [ ] Structured logging with levels (e.g., logrus, zap)
- [ ] Unit and integration tests
- [ ] Database integration (PostgreSQL, MongoDB)
- [ ] Move validation endpoints
- [ ] Game history API
- [ ] API rate limiting
- [ ] Metrics and health monitoring (Prometheus)
- [ ] Authentication/authorization (JWT)
- [ ] Redis caching layer

## 📚 Resources

- [Go HTTP Server Documentation](https://pkg.go.dev/net/http)
- [Go Best Practices](https://golang.org/doc/effective_go)
- [REST API Design Guidelines](https://restfulapi.net/)

## 🤝 Contributing

When adding new features to the backend:

1. Follow Go naming conventions
2. Add appropriate error handling
3. Include JSON request/response examples
4. Test endpoints manually with curl
5. Update this README with new endpoints