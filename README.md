# RSAG - RSS Aggregator API

**RSAG** is a production-ready backend API service written in Go for aggregating RSS feeds, managing users, and delivering feed content. RSAG is designed for scalability, security, and extensibility—making it ideal for building custom RSS readers, content aggregators, or for powering backend services requiring fresh feed data.

---

## Features

- **User Management:** Register new users, authenticate, and retrieve user data.
- **Feed Management:** Add and discover RSS feeds.
- **Feed Follows:** Users can follow/unfollow any feed.
- **Post Aggregation:** Periodically scrapes and stores new posts from followed feeds.
- **API Key Authentication:** All user-specific endpoints are protected by API keys.
- **Concurrent Scraping:** Efficient feed scraping using Go’s goroutines and worker pools.
- **CORS Support:** Easily integrate with web frontends via configurable CORS.
- **Environment-based Config:** Secure and portable configuration with `.env`.
- **PostgreSQL Storage:** Reliable, scalable relational database backend.

---

## API Overview

All endpoints are versioned under `/v1`:

| Method | Endpoint                               | Description                                  | Auth Required |
|--------|----------------------------------------|----------------------------------------------|--------------|
| GET    | `/v1/healthz`                          | Health check                                 | No           |
| GET    | `/v1/err`                              | Error test endpoint                          | No           |
| POST   | `/v1/users`                            | Create a new user                            | No           |
| GET    | `/v1/users`                            | Get current user (API key)                   | Yes          |
| POST   | `/v1/feeds`                            | Add a new feed                               | Yes          |
| GET    | `/v1/feeds`                            | List all feeds                               | No           |
| POST   | `/v1/feed_follows`                     | Follow a feed                                | Yes          |
| GET    | `/v1/feed_follows`                     | List followed feeds                          | Yes          |
| DELETE | `/v1/feed_follows/{feedFollowID}`      | Unfollow a feed                              | Yes          |
| GET    | `/v1/posts`                            | Get posts for the authenticated user         | Yes          |

---

## Quick Start

### 1. Clone the Repository

```sh
git clone https://github.com/asm2212/rsag.git
cd rsag
```

### 2. Setup Environment

Create a `.env` file in the project root:

```env
PORT=8080
DB_URL=postgres://user:password@localhost:5432/rsag?sslmode=disable
```

### 3. Database Setup

- Ensure PostgreSQL is running and accessible.
- Create a database for RSAG.
- Run database migrations in `/internal/database` if provided.

### 4. Run the Server

```sh
go run main.go
```

Server will start on the port specified in `.env`.

---

## Scraping & Aggregation

- The scraper runs in the background, fetching new posts from feeds every minute (configurable).
- Scraping is concurrent and uses a worker pool for efficiency.
- Duplicate posts are automatically skipped.

---

## Example Usage

1. **Create a User**

   ```bash
   curl -X POST http://localhost:8080/v1/users -d '{"name":"Alice"}'
   # Response includes your API key
   ```

2. **Add a Feed**

   ```bash
   curl -X POST http://localhost:8080/v1/feeds \
     -H "Authorization: ApiKey {YOUR_API_KEY}" \
     -d '{"name":"Go Blog","url":"https://blog.golang.org/feed.atom"}'
   ```

3. **Follow a Feed**

   ```bash
   curl -X POST http://localhost:8080/v1/feed_follows \
     -H "Authorization: ApiKey {YOUR_API_KEY}" \
     -d '{"feed_id":"<feed-uuid>"}'
   ```

4. **Fetch Posts**

   ```bash
   curl -X GET http://localhost:8080/v1/posts \
     -H "Authorization: ApiKey {YOUR_API_KEY}"
   ```

---

## Project Structure

```
.
├── main.go                 # Entry point, server setup, route registration
├── scraper.go              # RSS feed scraping logic and concurrency
├── types.go                # API types, model/data mapping
├── internal/
│   └── database/           # SQL queries, migrations, DB adapters
├── go.mod
├── go.sum
└── README.md
```

---

## Deployment

- **Production Readiness:** RSAG is designed for containerization and cloud deployment.
- **Port & DB config:** All critical configs are via environment variables.
- **Logging:** Errors and key events are logged to stdout/stderr.
- **Security:** API key authentication for all user data and write endpoints.

---

## Dependencies

- [Go](https://golang.org/)
- [go-chi/chi](https://github.com/go-chi/chi) - HTTP router
- [lib/pq](https://github.com/lib/pq) - PostgreSQL driver
- [joho/godotenv](https://github.com/joho/godotenv) - .env loader
- [google/uuid](https://github.com/google/uuid) - UUID generation

---

## Contributing

Contributions, bug reports, and feature requests are welcome!  
Please open an issue or submit a pull request.

---

## License

MIT License © 2025 asm2212

---

## Author

[asm2212](https://github.com/asm2212)
