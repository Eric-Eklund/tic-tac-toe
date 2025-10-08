# React Tic-Tac-Toe

A classic Tic-Tac-Toe game built with React, TypeScript, and Vite. Features a beautiful, responsive design with smooth animations and modern styling.

## Features

- 🎮 Classic 3x3 Tic-Tac-Toe gameplay
- 👥 Two-player mode (X vs O)
- 🎨 Beautiful gradient background and modern UI design
- ✨ Smooth animations and hover effects
- 📱 Responsive design that works on all devices
- ⚡ Built with Vite for fast development and optimized builds
- 🔧 TypeScript for type safety and better developer experience

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

1. The game starts with Player 1 (X) making the first move
2. Players take turns clicking on empty squares to place their symbol
3. The first player to get three of their symbols in a row (horizontally, vertically, or diagonally) wins
4. If all squares are filled without a winner, the game ends in a draw
5. Click "Rematch" to start a new game

## Project Structure

```
src/
├── components/          # React components
│   └── Player.tsx      # Player information component
├── App.tsx             # Main application component
├── main.tsx           # Application entry point
├── index.css          # Global styles and animations
└── App.css            # Component-specific styles
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the game.

## License

This project is open source and available under the [MIT License](LICENSE).
