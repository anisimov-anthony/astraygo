## About

**AstrayGo** is a play on the phrase "go astray" — because the whole point is to make sure your objects don't go astray. The service provides efficient storage and retrieval of geolocation data with spatial query capabilities.

Built with Go, AstrayGo leverages PostgreSQL with PostGIS extension for spatial data operations and Redis for high-speed caching, ensuring low-latency responses for location-based queries.

## Tech Stack

### Backend
- **Go 1.25+** - Core application language
- **Gin** - HTTP web framework for routing and middleware

### Database & Caching
- **PostgreSQL 17** - Primary data store
- **PostGIS 3.5** - Spatial database extension for geographic objects
- **Redis 8** - In-memory cache for fast object lookups

### Infrastructure
- **Docker & Docker Compose** - Containerization and orchestration
- **pgx/v5** - PostgreSQL driver for Go
- **go-redis/v9** - Redis client for Go

### Observability
- **ELK Stack** - Centralized logging and analysis (Elasticsearch, Logstash, Kibana)
- **Filebeat 8.12** - Log shipping from containers to ELK
- **Prometheus exporters** - Metrics collection (Node, PostgreSQL, Redis)
- **cAdvisor** - Container metrics
- **Zap** - Structured logging within the application

## Quick Start

```bash
docker compose up -d --build
```

The service will be available at `http://localhost:8080`

## License

MIT License - see [LICENSE](LICENSE)
