# Tic-Tac-Toe Full Stack Application

A modern, full-stack tic-tac-toe game featuring a beautiful React frontend with TypeScript and a Go backend server with real-time WebSocket support. The application is containerized with Docker for easy deployment and development.

## ✨ Features

- 🎮 Classic 3x3 Tic-Tac-Toe gameplay
- 🎨 Beautiful dark theme with orange glowing effects
- 👥 Two-player mode with editable player names
- 🏆 Game over screen with winner announcements
- 📊 Complete game move history
- 🔄 Backend-integrated game state management
- 🌐 WebSocket support for real-time features
- 💾 Persistent game storage to JSON file
- ⚡ Fast development with Vite and hot reloading
- 🐳 Dockerized for easy deployment
- 🔧 Full TypeScript support
- 📱 Responsive design

## 🏗️ Architecture

This is a full-stack application with:

- **Frontend**: React 19 + TypeScript + Vite + Axios
- **Backend**: Go 1.25 + Gin framework + Gorilla WebSocket
- **Storage**: JSON file-based persistence with volume mounting
- **Deployment**: Multi-stage Docker build for production

## 📁 Project Structure

```
tic-tac-toe/
├── back-end/              # Go backend server
│   ├── internal/          # Internal packages
│   │   ├── config/        # Configuration management
│   │   ├── handlers/      # HTTP request handlers
│   │   ├── models/        # Data models (Match, Player, GameBoard)
│   │   ├── services/      # Business logic and persistence
│   │   └── websocket/     # WebSocket hub and client management
│   ├── main.go            # Application entry point
│   ├── go.mod             # Go module dependencies
│   ├── matches.json       # Local game storage
│   └── README.md          # Backend-specific documentation
├── front-end/             # React frontend application
│   ├── src/               # React components and game logic
│   │   ├── components/    # UI components
│   │   ├── services/      # Backend API integration (Axios)
│   │   ├── types/         # TypeScript type definitions
│   │   └── assets/        # Game assets and data
│   ├── public/            # Static assets (images, etc.)
│   ├── package.json       # Node.js dependencies
│   ├── vite.config.ts     # Vite build configuration
│   └── README.md          # Frontend-specific documentation
├── data/                  # Persistent data directory (Docker volume)
│   └── matches.json       # Production game storage
├── Dockerfile             # Multi-stage Docker build
├── docker-compose.yml     # Docker composition
└── README.md              # This file
```

## 🚀 Quick Start

### Option 1: Docker (Recommended)

```bash
# Clone the repository
git clone <your-repo-url>
cd tic-tac-toe

# Build and run with Docker Compose
docker-compose up --build

# Or build and run manually
docker build -t tic-tac-toe .
docker run -p 8080:8080 tic-tac-toe
```

The application will be available at **http://localhost:8080**

### Option 2: Development Mode

**Prerequisites:**
- Node.js 18+ 
- Go 1.21+

**Run Frontend (React):**
```bash
cd front-end
npm install
npm run dev
# Available at http://localhost:5173
```

**Run Backend (Go):**
```bash
cd back-end
go mod download
go run main.go
# API and WebSocket available at http://localhost:8080
```

## 🔌 API Endpoints

### REST API
- `GET /api/health` - Health check endpoint
- `GET /api/new-match` - Start a new game match (returns match ID, initial board, and players)
- `GET /` - Serves the React application
- `GET /<any-path>` - Serves static assets (CSS, JS, images)

### WebSocket
- `WS /ws` - WebSocket connection for real-time game updates
  - Supports message types: `hello`, `welcome`, and custom game events
  - Broadcasts messages to all connected clients

## 🐳 Docker Architecture

The Dockerfile uses a multi-stage build process:

1. **Frontend Build Stage**: 
   - Uses `node:25-alpine`
   - Installs dependencies and builds React app with Vite
   - Outputs optimized static files to `dist/`

2. **Backend Build Stage**:
   - Uses `golang:1.25-alpine` 
   - Downloads Go modules and dependencies
   - Builds the Go HTTP server binary with build cache

3. **Production Stage**:
   - Uses minimal `alpine:latest`
   - Copies Go binary and React static files
   - Single container with both frontend and backend
   - Exposes port 8080
   - Mounts `./data` volume for persistent storage

## 🎮 How to Play

1. Open http://localhost:8080 in your browser
2. Click "Edit" to customize player names
3. Players take turns clicking empty squares
4. Get three in a row to win!
5. Click "Rematch!" to play again

## 🛠️ Development

### Backend Structure

The backend follows a clean architecture pattern:
- **handlers**: HTTP request handlers using Gin framework
- **models**: Data structures (Match, Player, GameBoard)
- **services**: Business logic and JSON file persistence
- **config**: Environment-based configuration with CORS
- **websocket**: Real-time communication hub

### Adding Backend API Endpoints

To add new API endpoints, edit `back-end/internal/handlers/handler.go`:

```go
func (h *Handler) SetupRoutes() *gin.Engine {
    router := gin.Default()
    // Add new routes
    router.GET("/api/your-endpoint", h.YourHandler)
    return router
}
```

### Frontend API Integration

Backend integration is handled via Axios in `front-end/src/services/backend.tsx`:

```typescript
import axios from "axios"

const api = axios.create({
    baseURL: "http://localhost:8080/api",
    timeout: 5000,
});

export async function yourApiCall(): Promise<YourType> {
    const response = await api.get<YourType>("/your-endpoint");
    return response.data;
}
```

## 📚 Documentation

- [Backend README](back-end/README.md) - Go server details
- [Frontend README](front-end/README.md) - React app details

## 🎯 Future Roadmap

- [x] Backend API integration
- [x] Game state persistence
- [x] WebSocket infrastructure
- [ ] Real-time multiplayer game synchronization
- [ ] Move validation on backend
- [ ] Game history API endpoints
- [ ] User authentication
- [ ] AI opponent
- [ ] Leaderboards
- [ ] Database integration (replace JSON files)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## 📄 License

This project is open source and available under the MIT License.
