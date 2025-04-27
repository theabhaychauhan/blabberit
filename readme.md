# BlabberIt 🔨

BlabberIt is a **self-hosted**, **peer-to-peer** chat application built in Go.  
It features **end-to-end encryption** (coming soon) and **offline message queuing** if a peer is unavailable.

---

## ✨ Features

- ✅ User registration with public key
- ✅ Peer-to-peer messaging (via upcoming libp2p)
- ✅ Offline message queuing (via local DB)
- ✅ PostgreSQL for persistent user and message data
- ✅ Fully containerized development (Docker for DB)
- ✅ CI/CD ready (GitHub Actions for testing)

---

## 🛠️ Tech Stack

- **Go** (1.21+)
- **PostgreSQL** (via Docker)
- **GORM** (ORM)
- **libp2p** (in progress)
- **Delve** (debugging)
- **Godotenv** (environment configs)

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/thechauhanabhay/blabberit.git
cd blabberit
```

### 2. Setup `.env.test`

```dotenv
POSTGRES_DSN=host=localhost user=blabberuser password=blabberpass dbname=blabberit_test port=5432 sslmode=disable TimeZone=UTC
```

### 3. Run the Server

```bash
go run ./cmd/server
```

Server will start at `http://localhost:8080`

### 4. Run Tests

```bash
go test ./internal/...
```

### 5. Debug Tests (optional)

```bash
dlv test ./internal/user
```

---

## 📦 API Endpoints (Current)

| Method | Path | Purpose |
|:-------|:-----|:--------|
| POST | `/register` | Register a new user with username + public key |
| GET | `/login?username=` | Retrieve public key by username |
| POST | `/send` | Send a message (HTTP relay fallback) |
| GET | `/inbox?user=` | Fetch messages for a public key |

---

## 🔥 Roadmap (Upcoming)

- [x] HTTP API first version
- [x] Offline message queuing
- [ ] Peer-to-peer messaging via **libp2p** (in progress)
- [ ] End-to-end encryption for messages
- [ ] NAT traversal / relay support
- [ ] Minimal CLI Client
- [ ] Hybrid fallback (P2P + server relay)

---

## ⚖️ License

[MIT](LICENSE)

---

## 🚀 Summary

BlabberIt is ready with HTTP APIs, offline queuing, Postgres storage, tests, and CI/CD.  
Peer-to-peer messaging via **libp2p** is coming next to enable serverless chatting across the internet!

