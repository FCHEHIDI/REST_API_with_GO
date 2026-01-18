# REST API with Go and Gin Framework

![REST API Architecture](./assets/Copilot_20260118_232258.png)

> **Building scalable backend systems** — Learning Go, REST APIs, and production-grade architecture patterns.

---

## 🚀 Overview

Full-featured REST API demonstrating modern backend engineering with **Go and Gin framework**. Built with clean architecture and designed to evolve into a production-ready, scalable system.

### Current Stack
✅ **REST API** with CRUD operations  
✅ **JWT Authentication** & authorization  
✅ **SQLite Database** with prepared statements  
✅ **Password Hashing** (bcrypt)  
✅ **Middleware** for protected routes  
✅ **Clean Architecture** - models, routes, middleware separation

### Scaling Roadmap 🔄
PostgreSQL • Redis • Docker • Kubernetes • Microservices • CI/CD • Monitoring • Rate Limiting • Testing Suite

---

## 🏗️ Architecture

```
├── db/          # Database layer
├── middlewares/ # Authentication & middleware
├── models/      # Data models & DB operations
├── routes/      # HTTP handlers & endpoints
├── utils/       # JWT & hashing utilities
└── api-test/    # API test requests
```

---
├── api-test/              # HTTP test files for REST Client
├── main.go                # Application entry point
├── go.mod                 # Go module definition
└── api.db                 # SQLite database (generated)
```

## � API Endpoints

### Public
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Welcome message |
| GET | `/events` | List all events |
| GET | `/events/:id` | Get event details |
| POST | `/signup` | Create account |
| POST | `/login` | Login (returns JWT) |

### Protected (Requires JWT)
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/events` | Create event |
| PUT | `/events/:id` | Update event |
| DELETE | `/events/:id` | Delete event |
| POST | `/events/:id/register` | Register for event |
| DELETE | `/events/:id/unregister` | Unregister from event |

---

## ⚡ Quick Start

```bash
# Clone and setup
git clone https://github.com/FCHEHIDI/REST_API_with_GO.git
cd REST_API_with_GO
go mod download

# Run
go run .
# Server at http://localhost:8080
```

### Test Flow
1. **Signup**: `POST /signup` → email & password
2. **Login**: `POST /login` → get JWT token
3. **Create Event**: `POST /events` + `Authorization: <token>`
4. Use test files in `api-test/` folder

---

## 📖 What I Learned

### Go Language Concepts
- Structs, interfaces, and methods
- Error handling patterns
- Package management with `go mod`
- Pointer vs value receivers
- Goroutines and concurrency concepts (for future enhancement)

### Web Development
- HTTP request/response cycle
- RESTful API design principles
- Status codes and proper error responses
- Content negotiation and JSON handling
- Middleware pattern for cross-cutting concerns

### Database Operations
- SQL query writing
- Prepared statements for security
- Database connection pooling
- Schema design and foreign keys
- Transaction management concepts

### Security
- Password hashing with bcrypt
- JWT token generation and validation
- Authentication vs Authorization
- Protecting routes and resources
- Security best practices

### Software Architecture
- Separation of concerns (MVC-like pattern)
- Handler → Model → Database flow
- Middleware for authentication
- Utility functions for reusability
- Clean code principles

## 🔄 Roadmap & Future Enhancements

This project is actively evolving. Here's what's coming next:

### Phase 2: Performance & Scalability
- [ ] **PostgreSQL Migration** - Production-grade relational database
- [ ] **Redis Caching** - Performance optimization with caching layer
- [ ] **Connection Pooling** - Efficient database connection management
- [ ] **Query Optimization** - Indexes and optimized queries

### Phase 3: Architecture & DevOps
- [ ] **Docker Containerization** - Containerize application and dependencies
- [ ] **Kubernetes Deployment** - Orchestration for scaling
- [ ] **CI/CD Pipeline** - GitHub Actions for automated testing and deployment
- [ ] **Microservices** - Break into independent services (auth, events, notifications)

### Phase 4: Advanced Features
- [ ] **GraphQL API** - Alternative to REST for flexible queries
- [ ] **WebSocket Support** - Real-time event updates
- [ ] **File Upload** - Event image/document handling
- [ ] **Email Notifications** - Event reminders and confirmations
- [ ] **Search Engine** - Elasticsearch integration for advanced search
- [ ] **Pagination & Filtering** - Efficient data retrieval
- [ ] **API Versioning** - Support multiple API versions

### Phase 5: Production Readiness
- [ ] **Rate Limiting** - Protect against abuse
- [ ] **API Monitoring** - Prometheus & Grafana
- [ ] **Logging System** - Structured logging with ELK stack
- [ ] **Comprehensive Testing** - Unit, integration, and E2E tests
- [ ] **API Documentation** - Swagger/OpenAPI specs
- [ ] **Security Hardening** - HTTPS, CORS, security headers
- [ ] **Load Testing** - Performance benchmarking

### Phase 6: Cloud & Distribution
- [ ] **Cloud Deployment** - AWS/GCP/Azure deployment
- [ ] **CDN Integration** - Global content delivery
- [ ] **Multi-region Support** - Geographic distribution
- [ ] **Backup & Disaster Recovery** - Automated backups

## 💡 Why This Matters

This isn't just a learning project—it's a **foundation for building real-world, scalable systems**. Each enhancement teaches critical production engineering skills that translate directly to professional backend development.

---

## 🛠️ Tech Stack

**Go** • **Gin Framework** • **SQLite** (→ PostgreSQL) • **JWT** • **bcrypt**

---

## 📚 What I'm Learning

Backend fundamentals • REST API design • Authentication patterns • Database operations • Clean architecture • Security best practices • Scalability planning

---

**Status**: ✅ Phase 1 Complete | 🔄 Ready to Scale | 🚀 Production-Bound

---

*Building this project to master backend engineering. Each phase adds production-grade features and real-world patterns.*

- Go documentation and community
- Gin framework documentation
- Various backend engineering tutorials and resources
- Open source community for excellent libraries

## 📬 Contact

Created as part of my journey to become a professional backend engineer. Connect with me to discuss backend development, Go, or software engineering!

---

**Note**: This is a learning project. While it implements production-ready patterns, it's designed for educational purposes. Always follow your organization's specific security and architectural guidelines for production applications.
