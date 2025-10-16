# React Tic-Tac-Toe

A modern, fully-featured Tic-Tac-Toe game built with React, TypeScript, and Vite. Features a stunning dark theme with orange glowing effects, smooth animations, and a complete game experience with player management and game history.

## Features

- 🎮 Classic 3x3 Tic-Tac-Toe gameplay with smart winner detection
- 👥 Two-player mode with editable player names
- 🏆 Game over screen with winner announcements and rematch functionality
- 📊 Game log showing move history
- 🎨 Beautiful dark theme with orange glowing effects and animations
- ✨ Epic title animations and smooth hover effects
- 🔧 Fully typed with TypeScript and structured Player objects
- 📱 Responsive design that works on all devices
- ⚡ Built with Vite for fast development and optimized builds
- 🎯 Proper game state management with turn-based logic

## Tech Stack

- **React** - UI library for building interactive components
- **TypeScript** - Type-safe JavaScript for better code quality
- **Vite** - Fast build tool and development server
- **CSS3** - Modern styling with gradients, animations, and flexbox
- **ESLint** - Code linting for consistent code style

## Getting Started

### Prerequisites

- Node.js (version 18 or higher)
- npm or yarn

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Eric-Eklund/tic-tac-toe.git
   cd tic-tac-toe
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```

4. Open your browser and navigate to `http://localhost:5173`

## Available Scripts

- `npm run dev` - Start the development server
- `npm run build` - Build the project for production
- `npm run preview` - Preview the production build locally
- `npm run lint` - Run ESLint to check code quality

## How to Play

1. **Edit Player Names**: Click the "Edit" button next to any player to customize their name
2. **Make Moves**: Players take turns clicking on empty squares to place their symbol (X or O)
3. **Active Player**: The current player is highlighted with an orange glow
4. **Win Condition**: Get three symbols in a row (horizontally, vertically, or diagonally) to win
5. **Draw**: If all squares are filled without a winner, the game ends in a draw
6. **Game Over**: A beautiful overlay appears showing the winner or draw result
7. **Rematch**: Click "Rematch!" to start a new game with the same players
8. **Game Log**: View the complete move history at the bottom of the screen

## Project Structure

```
src/
├── components/          # React components
│   ├── GameBoard.tsx   # Game board with 3x3 grid
│   ├── GameOver.tsx    # Game over overlay component
│   ├── Log.tsx         # Game move history component
│   └── Player.tsx      # Player information and name editing
├── types/              # TypeScript type definitions
│   └── shared.types.tsx # Shared types (Player, GameTurn, etc.)
├── assets/             # Game assets and data
│   └── winning-combinations.ts # All possible winning combinations
├── App.tsx             # Main application component with game logic
├── main.tsx           # Application entry point
└── index.css          # Complete styling with dark theme and animations
```

## Architecture

The game uses a clean, component-based architecture with:

- **Structured Player Objects**: Each player has an ID, name, and symbol
- **Type Safety**: Full TypeScript coverage with custom types
- **State Management**: React hooks for game state and player management
- **Smart Game Logic**: Automatic winner detection and turn management
- **Component Separation**: Each UI element is a focused, reusable component

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the game.

## License

This project is open source and available under the [MIT License](LICENSE).
