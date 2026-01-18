# 🚀 Quick Start

## Setup & Run

```bash
# Install dependencies
go mod download

# Run the server
go run .

# Server starts at http://localhost:8080
```

## Test Flow

1. **Signup**: `POST /signup` with email & password
2. **Login**: `POST /login` - get your JWT token
3. **Use Token**: Add `Authorization: YOUR_TOKEN` header for protected routes
4. **Create Event**: `POST /events` (requires token)
5. **Manage Events**: GET, PUT, DELETE endpoints available

Check `api-test/*.http` files for ready-to-use requests!

## Ready to Commit

```bash
git init
git add .
git commit -m "Initial commit: Go REST API with JWT auth"
git remote add origin <your-repo-url>
git push -u origin main
```

---

**Next**: Scale up with PostgreSQL, Docker, Redis, and microservices! 🚀
