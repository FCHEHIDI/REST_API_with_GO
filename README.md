# REST API with Go and Gin Framework

> **Learning Project**: Building a scalable REST API with Go to master backend engineering and prepare for production-grade systems.

## 🚀 Project Overview

A hands-on REST API built with Go, demonstrating core backend concepts with plans to evolve into a production-ready, scalable system. This project serves as a foundation for learning modern API development while maintaining clean, extensible architecture.

**Current Features:**
- ✅ RESTful API with CRUD operations
- ✅ JWT authentication & authorization
- ✅ SQLite database integration
- ✅ Secure password hashing
- ✅ Protected routes with middleware
- ✅ Clean code architecture

**Future Vision:**
- 🔄 Migration to PostgreSQL for production
- 🔄 Microservices architecture
- 🔄 Redis caching layer
- 🔄 Docker containerization
- 🔄 Kubernetes orchestration
- 🔄 Comprehensive testing suite
- 🔄 CI/CD pipeline
- 🔄 API rate limiting & monitoring
- 🔄 Advanced features (search, pagination, filtering)

## 🎯 Learning Objectives

Mastering backend engineering fundamentals while building a system designed for real-world scaling and enhancement.

## 🏗️ Project Structure

```
RESTAPI_GO/
├── db/                    # Database initialization and schema
│   └── db.go
├── middlewares/           # Custom middleware (authentication)
│   └── auth.go
├── models/                # Data models and database operations
│   ├── event.go
│   └── user.go
├── routes/                # HTTP handlers and route definitions
│   ├── events.go
│   ├── users.go
│   ├── register.go
│   └── routes.go
├── utils/                 # Utility functions (JWT, hashing)
│   ├── hash.go
│   └── jwt.go
├── api-test/              # HTTP test files for REST Client
├── main.go                # Application entry point
├── go.mod                 # Go module definition
└── api.db                 # SQLite database (generated)
```

## 🚀 Features

### Event Management
- ✅ Create, read, update, and delete events
- ✅ Event registration and unregistration
- ✅ Authorization checks (only event creators can modify)

### User Management
- ✅ User signup with password hashing
- ✅ User login with JWT token generation
- ✅ Protected routes requiring authentication

### Security
- ✅ Password hashing with bcrypt
- ✅ JWT-based authentication
- ✅ Authorization middleware
- ✅ SQL injection prevention with prepared statements

## 📋 Prerequisites

- Go 1.24 or higher
- Git
- A code editor (VS Code recommended)
- REST Client (VS Code extension) or Postman

## ⚙️ Installation

1. **Clone the repository**
   ```bash
   git clone <your-repository-url>
   cd RESTAPI_GO
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the application**
   ```bash
   go run .
   ```

The server will start on `http://localhost:8080`

## 🔌 API Endpoints

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Welcome message |
| GET | `/events` | Get all events |
| GET | `/events/:id` | Get event by ID |
| POST | `/signup` | Register new user |
| POST | `/login` | Login and get JWT token |

### Protected Endpoints (Require Authentication)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/events` | Create new event |
| PUT | `/events/:id` | Update event (creator only) |
| DELETE | `/events/:id` | Delete event (creator only) |
| POST | `/events/:id/register` | Register for an event |
| DELETE | `/events/:id/unregister` | Unregister from an event |

## 🧪 Testing the API

### Using REST Client (VS Code Extension)

Test files are available in the `api-test/` directory:

1. Install the REST Client extension in VS Code
2. Open any `.http` file in `api-test/`
3. Click "Send Request" above each request

### Using cURL

**Sign up a new user:**
```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

**Login:**
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

**Create an event (requires token):**
```bash
curl -X POST http://localhost:8080/events \
  -H "Content-Type: application/json" \
  -H "Authorization: YOUR_JWT_TOKEN" \
  -d '{
    "name":"Tech Conference",
    "description":"Annual tech event",
    "location":"San Francisco",
    "dateTime":"2026-06-15T10:00:00Z"
  }'
```

**Get all events:**
```bash
curl http://localhost:8080/events
```

## 🔑 Authentication Flow

1. **Sign up**: Create a new user account with email and password
2. **Login**: Receive a JWT token (valid for 72 hours)
3. **Use Token**: Include token in `Authorization` header for protected routes
4. **Token Format**: `Authorization: YOUR_JWT_TOKEN_HERE`

## 💾 Database Schema

### Users Table
```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);
```

### Events Table
```sql
CREATE TABLE events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    dateTime DATETIME NOT NULL,
    userID INTEGER,
    FOREIGN KEY(userID) REFERENCES users(id)
);
```

### Registrations Table
```sql
CREATE TABLE registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userID INTEGER,
    eventID INTEGER,
    FOREIGN KEY(userID) REFERENCES users(id),
    FOREIGN KEY(eventID) REFERENCES events(id)
);
```

## 🛠️ Technologies Used

- **[Go](https://golang.org/)** - Programming language
- **[Gin](https://github.com/gin-gonic/gin)** - Web framework
- **[modernc.org/sqlite](https://gitlab.com/cznic/sqlite)** - Pure Go SQLite driver
- **[golang-jwt/jwt](https://github.com/golang-jwt/jwt)** - JWT implementation
- **[golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto)** - Cryptography (bcrypt)

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

## 🤝 Contributing

This is a personal learning project, but suggestions and feedback are welcome! Feel free to:
- Open issues for bugs or improvements
- Suggest better practices or patterns
- Share learning resources

## 📝 License

This project is created for educational purposes. Feel free to use it for your own learning.

## 🙏 Acknowledgments

- Go documentation and community
- Gin framework documentation
- Various backend engineering tutorials and resources
- Open source community for excellent libraries

## 📬 Contact

Created as part of my journey to become a professional backend engineer. Connect with me to discuss backend development, Go, or software engineering!

---

**Note**: This is a learning project. While it implements production-ready patterns, it's designed for educational purposes. Always follow your organization's specific security and architectural guidelines for production applications.
