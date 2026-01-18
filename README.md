# REST API with Go and Gin Framework

![REST API Architecture](./assets/Copilot_20260118_232258.png)

> **Production-grade backend engineering** — Exploring Go ecosystem and modern scalable architecture patterns.

---

## 🚀 Overview

Professional REST API implementation showcasing **Go and Gin framework** best practices. Built with clean architecture and designed for progressive evolution into distributed, cloud-native systems.

### Technical Implementation
✅ **RESTful API** with comprehensive CRUD operations  
✅ **JWT Authentication** with secure authorization flow  
✅ **SQLite** → PostgreSQL migration-ready architecture  
✅ **bcrypt** password hashing with industry standards  
✅ **Middleware patterns** for cross-cutting concerns  
✅ **Clean Architecture** with clear separation of concerns

### Evolution Path 🔄
PostgreSQL • Redis • Docker • Kubernetes • Microservices • CI/CD • Monitoring • Rate Limiting • GraphQL

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

## � Architecture Evolution Roadmap

Progressive enhancement toward cloud-native, distributed systems:

**Phase 2**: PostgreSQL with advanced query optimization, Redis caching strategy, connection pooling  
**Phase 3**: Containerization (Docker), orchestration (Kubernetes), GitOps CI/CD, service mesh  
**Phase 4**: GraphQL federation, WebSocket real-time layer, distributed file storage, Elasticsearch  
**Phase 5**: Rate limiting algorithms, observability stack (Prometheus/Grafana/Jaeger), comprehensive testing  
**Phase 6**: Multi-cloud deployment, edge computing, disaster recovery automation

---

## 🛠️ Tech Stack

**Go 1.24** • **Gin Framework** • **SQLite** → **PostgreSQL** • **JWT** • **bcrypt** • **modernc.org/sqlite**

---

## 🎯 Technical Focus

Exploring Go ecosystem • Microservices patterns • Distributed systems • Cloud-native architecture • Performance optimization • Security hardening • DevOps practices

---

**Status**: ✅ Foundation Complete | 🔄 Actively Evolving | 🚀 Cloud-Native Ready

---

*Continuous exploration of backend technologies and architectural patterns. Each phase introduces production-grade patterns and modern distributed systems concepts.*


## 📬 Contact

Connect with me to discuss backend development, Go, or software engineering!

---

**Note**: This is a learning project. While it implements production-ready patterns, it's designed for educational purposes. Always follow your organization's specific security and architectural guidelines for production applications.
