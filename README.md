# Shrinkr

**Shrinkr** is a URL shortening service built using Go, with features for generating short URLs, QR generation and handling redirection. It also includes Prometheus monitoring for tracking usage metrics, Redis for caching url.

## Features
- Generate shortened URLs.
- Redirect to the original URL when accessed.
- Monitor and track performance with Prometheus.

## Tech Stack
- **Go** for backend server and logic.
- **Redis** for caching.
- **TypeScript** for the client-side web app.
- **Prometheus** for monitoring.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/hritesh04/shrinkr.git
   ```
2. Set up dependencies:
   ```bash
   go mod tidy
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

## Usage
1. Access the URL shortener via the client web app.
2. Generate QR code png using client web app.
2. Manage and analyze shortened URLs.
