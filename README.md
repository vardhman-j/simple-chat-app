# Realtime Chat App

A simple real-time chat application built with a Go WebSocket backend and a React frontend.

## Features

- Real-time messaging over WebSockets
- Multi-client chat pool on the backend
- React frontend UI for sending and displaying messages
- Simple local development setup

## Project Structure

```text
simple-chat-app/
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── pkg/
│       └── websocket/
│           ├── Client.go
│           ├── pool.go
│           └── websocket.go
├── frontend/
│   ├── package.json
│   ├── public/
│   └── src/
├── README.md
└── .gitignore
```

## Tech Stack

- Backend: Go
- Frontend: React + JavaScript
- Communication: Gorilla WebSocket

## Prerequisites

- Go 1.22 or newer
- Node.js 18+ and npm
- A modern browser

## Getting Started

### 1. Start the backend

From the project root:

```bash
cd backend
go mod tidy
go run main.go
```

The backend listens on:

```text
http://localhost:8080/ws
```

### 2. Start the frontend

Open a second terminal and run:

```bash
cd frontend
npm install
npm start
```

This will start the React app, usually at:

```text
http://localhost:3000
```

## Notes

- The frontend connects to the backend WebSocket endpoint at `ws://localhost:8080/ws`.
- The backend currently broadcasts messages to all connected clients in the connection pool.

## Development

If you want to extend the app:

- add message formatting or chat history features in the frontend
- expand the Go server with user names, rooms, or persistence
- improve UI styling and message layout

## License

This project is for learning and local development purposes.
