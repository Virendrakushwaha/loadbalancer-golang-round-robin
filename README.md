# Load Balancer - Round Robin (Go)

A simple HTTP load balancer implementation in Go using round-robin algorithm to distribute requests across multiple backend servers.

## Features

- Round-robin load balancing algorithm
- HTTP reverse proxy functionality
- Health check support (basic implementation)
- Configurable backend servers
- Test endpoint for debugging

## Project Structure

```
loadbalancer-golang-round-robin/
├── models/
│   ├── LoadBalancer.go    # Load balancer implementation
│   └── Server.go          # Server model and proxy setup
├── utils/
│   └── helpers.go         # Utility functions
├── main.go               # Entry point
├── go.mod               # Go module file
└── README.md           # This file
```

## Usage

### Running the Load Balancer

```bash
go run main.go
```

The load balancer will start on `localhost:8080`.

### Testing Round-Robin

**Main endpoint**: Visit `http://localhost:8080/` - redirects to backend servers

### Backend Servers

The load balancer is configured with three backend servers:
- https://www.github.com
- https://www.google.com
- https://www.facebook.com

## How It Works

1. **Round-Robin Algorithm**: Requests are distributed sequentially across all available servers
2. **Reverse Proxy**: Uses Go's `httputil.ReverseProxy` to forward requests
3. **Health Checks**: Basic health check implementation (currently always returns true)

## API Endpoints

- `GET /` - Load balanced requests to backend servers

## Configuration

To modify backend servers, edit the `servers` slice in `main.go`:

```go
servers := []models.Server{
    models.NewServer("https://your-server-1.com"),
    models.NewServer("https://your-server-2.com"),
    models.NewServer("https://your-server-3.com"),
}
```

