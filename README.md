# Tic-Tac-Toe Full Stack Application

A modern, full-stack tic-tac-toe game featuring a beautiful React frontend with TypeScript and a Go backend server. The application is containerized with Docker for easy deployment and development.

## ✨ Features

- 🎮 Classic 3x3 Tic-Tac-Toe gameplay
- 🎨 Beautiful dark theme with orange glowing effects
- 👥 Two-player mode with editable player names
- 🏆 Game over screen with winner announcements
- 📊 Complete game move history
- ⚡ Fast development with Vite and hot reloading
- 🐳 Dockerized for easy deployment
- 🔧 Full TypeScript support
- 📱 Responsive design

## 🏗️ Architecture

This is a full-stack application with:

- **Frontend**: React + TypeScript + Vite (runs in browser)
- **Backend**: Go HTTP server (serves static files + API endpoints)
- **Deployment**: Multi-stage Docker build for production

## 📁 Project Structure

```
tic-tac-toe/
├── back-end/              # Go backend server
│   ├── main.go           # HTTP server with static file serving
│   ├── go.mod            # Go module dependencies
│   └── README.md         # Backend-specific documentation
├── front-end/            # React frontend application
│   ├── src/              # React components and game logic
│   ├── public/           # Static assets (images, etc.)
│   ├── package.json      # Node.js dependencies
│   ├── vite.config.ts    # Vite build configuration
│   └── README.md         # Frontend-specific documentation
├── Dockerfile            # Multi-stage Docker build
├── docker-compose.yml    # Docker composition
└── README.md            # This file
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
go run main.go
# API available at http://localhost:8080
```

## 🔌 Current API Endpoints

- `GET /api/health` - Health check endpoint
- `GET /` - Serves the React application
- `GET /<any-path>` - Serves static assets (CSS, JS, images)

> **Note**: Game logic currently runs entirely in the frontend. Backend API endpoints for game state management will be added in future iterations.

## 🐳 Docker Architecture

The Dockerfile uses a multi-stage build process:

1. **Frontend Build Stage**: 
   - Uses `node:18-alpine`
   - Installs dependencies and builds React app with Vite
   - Outputs optimized static files to `dist/`

2. **Backend Build Stage**:
   - Uses `golang:1.21-alpine` 
   - Builds the Go HTTP server binary

3. **Production Stage**:
   - Uses minimal `alpine:latest`
   - Copies Go binary and React static files
   - Single container with both frontend and backend

## 🎮 How to Play

1. Open http://localhost:8080 in your browser
2. Click "Edit" to customize player names
3. Players take turns clicking empty squares
4. Get three in a row to win!
5. Click "Rematch!" to play again

## 🛠️ Development

### Adding Backend API Endpoints

To add new API endpoints, edit `back-end/main.go`:

```go
// Add after line 14
http.HandleFunc("/api/game/start", startGameHandler)
http.HandleFunc("/api/game/move", makeMoveHandler)
```

### Frontend API Integration

To call backend APIs from React:

```js
// Example API call
const response = await fetch('/api/game/start', { method: 'POST' });
const data = await response.json();
```

## 📚 Documentation

- [Backend README](back-end/README.md) - Go server details
- [Frontend README](front-end/README.md) - React app details

## 🎯 Future Roadmap

- [ ] Move game logic to backend
- [ ] Add multiplayer support
- [ ] WebSocket real-time updates
- [ ] Game history persistence
- [ ] User authentication
- [ ] AI opponent
- [ ] Leaderboards

## 🤝 Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## 📄 License

This project is open source and available under the MIT License.
