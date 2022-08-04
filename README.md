# Chat App with WebSockets in Go and React.js
Real-time chat app using React, Go, Websocket, Redis(Pub/Sub).

## Launch
```bash
make up
# or
docker-compose up --build
```

## Directory structure
```
.
├── docker-compose.yaml
├── Makefile
├── client
│   ├── Dockerfile
│   ├── package.json
│   ├── package-lock.json
│   ├── vite.config.ts
│   ├── index.html
│   └── src
│       ├── App.tsx
│       ├── main.tsx
│       ├── components
│       ├── hooks
│       ├── models
│       └── state
└── server
    ├── go.mod
    ├── go.sum
    ├── Dockerfile
    ├── .air.toml
    └── src
        ├── main.go
        ├── services
        ├── domain
        └── handlers
```