
# Conway's Game of Life Web App

## Project Structure

Create a monorepo with separate backend and frontend directories:

```
gameOfLife/
├── backend/
│   ├── cmd/server/main.go          # HTTP server entry point
│   ├── internal/
│   │   ├── game/                   # Game logic
│   │   │   ├── grid.go            # Grid data structure
│   │   │   ├── rules.go           # Conway's rules
│   │   │   └── simulator.go       # Game loop
│   │   └── api/                    # HTTP handlers
│   │       └── handlers.go        # REST endpoints
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   │   ├── Grid.tsx           # Canvas grid component
│   │   │   └── Controls.tsx       # Play/pause/step/reset
│   │   ├── services/
│   │   │   └── api.ts             # Backend API client
│   │   ├── App.tsx                # Main app
│   │   └── index.tsx              # Entry point
│   ├── public/
│   │   └── index.html
│   ├── package.json
│   └── tsconfig.json
└── README.md
```

## Backend Implementation (Go)

### 1. Game Logic (`backend/internal/game/`)

**grid.go**: Implement grid data structure with methods to get/set cell states and neighbor counting.

**rules.go**: Implement Conway's rules:

- Live cell: 2-3 neighbors → survives, else dies
- Dead cell: exactly 3 neighbors → becomes alive

**simulator.go**: Handle game state transitions and provide methods to step forward one generation.

### 2. REST API (`backend/internal/api/handlers.go`)

Endpoints:

- `POST /api/game/init` - Initialize new game with dimensions
- `POST /api/game/step` - Advance one generation
- `GET /api/game/state` - Get current grid state
- `POST /api/game/toggle` - Toggle cell at position
- `POST /api/game/reset` - Clear grid

### 3. Server (`backend/cmd/server/main.go`)

- Set up HTTP router
- Serve static frontend files from `frontend/build`
- Enable CORS for development
- Listen on port 8080

## Frontend Implementation (React + TypeScript)

### 1. Grid Component (`Grid.tsx`)

- Render grid using HTML5 Canvas 2D
- Handle mouse clicks to toggle cells
- Update canvas when game state changes
- Draw cells as filled rectangles

### 2. Controls Component (`Controls.tsx`)

Buttons:

- Play/Pause (toggle simulation)
- Step (advance one generation)
- Reset (clear grid)
- Speed control (optional interval slider)

### 3. API Service (`services/api.ts`)

- Fetch wrapper functions for all backend endpoints
- Handle JSON serialization
- Error handling

### 4. Main App (`App.tsx`)

- State management for game grid, playing status
- Game loop using `setInterval` when playing
- Layout grid and controls
- Initialize game on mount

## Setup Steps

1. Clean up old GopherJS files (main.go, index.html)
2. Initialize backend Go module in `backend/`
3. Initialize React + TypeScript app in `frontend/` using Vite or Create React App
4. Implement backend game logic and API
5. Implement frontend components and wire up API calls
6. Update root README with build/run instructions

### To-dos

- [ ] Remove old GopherJS files (main.go, index.html, go.mod, go.sum)
- [ ] Create backend directory structure and initialize Go module
- [ ] Implement Conway's rules in grid.go, rules.go, and simulator.go
- [ ] Create REST API handlers and HTTP server
- [ ] Initialize React + TypeScript app with Vite
- [ ] Build Canvas-based Grid component with click handling
- [ ] Build Controls component with play/pause/step/reset buttons
- [ ] Create API service and connect frontend to backend
- [ ] Wire up Grid and Controls in App.tsx with game loop
- [ ] Update README with build and run instructions