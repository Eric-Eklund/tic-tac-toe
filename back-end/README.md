# Tic-Tac-Toe Backend Server

A lightweight Go HTTP server that serves the React frontend and provides API endpoints for the tic-tac-toe game.

## 🏗️ Architecture

This Go backend serves dual purposes:

1. **Static File Server**: Serves the built React application and its assets
2. **API Server**: Provides RESTful endpoints for game logic (future expansion)

## 🔧 Tech Stack

- **Go 1.21+** - Backend language
- **Go Standard Library** - HTTP server (`net/http`)
- **Alpine Linux** - Production container base image

## 📁 Project Structure

```
back-end/
├── main.go       # HTTP server with routing and handlers
├── go.mod        # Go module definition
├── go.sum        # Go module checksums
└── README.md     # This file
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Port 8080 available

### Development

```bash
# Navigate to backend directory
cd back-end

# Run the development server
go run main.go

# Server starts at http://localhost:8080
```

### Build Binary

```bash
# Build for current platform
go build -o tic-tac-toe-server main.go

# Build for Linux (production)
CGO_ENABLED=0 GOOS=linux go build -o main .
```

## 🔌 API Endpoints

### Current Endpoints

| Method | Path | Description | Response |
|--------|------|-------------|----------|
| GET | `/api/health` | Health check endpoint | `{"status":"ok","service":"tic-tac-toe-backend"}` |
| GET | `/` | Serves React application | HTML page |
| GET | `/*` | Serves static assets | CSS, JS, images, etc. |

### Future API Endpoints (Planned)

```go
// Game management
POST   /api/game/start     // Start a new game
GET    /api/game/state     // Get current game state
POST   /api/game/move      // Make a move
DELETE /api/game/reset     // Reset/restart game

// Player management
POST   /api/players        // Update player names
GET    /api/players        // Get player information

// Game history
GET    /api/history        // Get game move history
POST   /api/history/clear  // Clear game history
```

## 📂 Static File Serving

The server uses Go's built-in `http.FileServer` to serve static files:

```go
fs := http.FileServer(http.Dir("./static"))
http.Handle("/", fs)
```

**File Structure in Production:**
```
/app/
├── main           # Go binary
└── static/        # React build output
    ├── index.html
    ├── assets/
    │   ├── index-[hash].css
    │   └── index-[hash].js
    ├── bg-pattern-dark.png
    └── vite.svg
```

## 🔧 Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP server port (future enhancement) |
| `STATIC_DIR` | `./static` | Static files directory (future enhancement) |

### Server Configuration

Currently hardcoded in `main.go`:
- **Port**: `:8080`
- **Static Directory**: `./static`
- **CORS**: Not configured (same-origin by default)

## 🛠️ Development

### Adding New API Endpoints

1. **Add route handler in `main.go`:**
```go
func main() {
    // Add new route
    http.HandleFunc("/api/game/start", startGameHandler)
    
    // Existing code...
}
```

2. **Implement handler function:**
```go
func startGameHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, `{"status":"game_started","id":"abc123"}`)
}
```

### Error Handling Best Practices

```go
func gameHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    // Handle request...
    
    if err != nil {
        log.Printf("Error: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
}
```

### JSON Request/Response Handling

```go
func jsonHandler(w http.ResponseWriter, r *http.Request) {
    // Parse JSON request
    var req struct {
        PlayerName string `json:"player_name"`
        Position   int    `json:"position"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    // Send JSON response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": true,
        "message": "Move processed",
    })
}
```

## 🐳 Docker Integration

The backend is built in a multi-stage Docker build:

**Build Stage:**
```dockerfile
FROM golang:1.21-alpine AS backend-build
WORKDIR /app/backend
COPY back-end/go.* ./
RUN go mod download
COPY back-end/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
```

**Production Stage:**
```dockerfile
COPY --from=backend-build /app/backend/main .
CMD ["./main"]
```

## 🧪 Testing

### Manual Testing

```bash
# Test health endpoint
curl http://localhost:8080/api/health

# Test static file serving
curl http://localhost:8080/

# Test asset serving
curl http://localhost:8080/bg-pattern-dark.png
```

### Future Unit Testing

```go
// Example test structure
func TestHealthHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/api/health", nil)
    w := httptest.NewRecorder()
    
    healthHandler(w, req)
    
    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }
}
```

## 📊 Performance Considerations

- **Static File Caching**: Go's FileServer includes basic caching headers
- **Binary Size**: ~6-8MB statically compiled binary
- **Memory Usage**: Minimal (~10-20MB at runtime)
- **Concurrency**: Go's HTTP server handles concurrent requests efficiently

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

- [ ] Environment-based configuration
- [ ] Structured logging with levels
- [ ] Request/response middleware
- [ ] Database integration
- [ ] WebSocket support for real-time features
- [ ] Unit and integration tests
- [ ] API rate limiting
- [ ] CORS configuration
- [ ] Metrics and health monitoring

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