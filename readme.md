# BlabberIt 🗨️

BlabberIt is a self-hosted, peer-to-peer chat application built in Go.  
It features end-to-end encryption and stores messages locally when the recipient is offline.

## Features

- ✅ User registration with public key
- ✅ End-to-end encryption
- ✅ Peer-to-peer messaging
- ✅ Offline message queuing
- ✅ PostgreSQL for persistent user data

## Tech Stack

- Go
- PostgreSQL
- GORM
- Docker (for DB)
- Delve (for debugging)
- Godotenv (for env config)

## Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/thechauhanabhay/blabberit.git
cd blabberit
```

### 2. Setup `.env.test`

```env
POSTGRES_DSN=host=localhost user=blabberuser password=blabberpass dbname=blabberit_test port=5432 sslmode=disable TimeZone=UTC
```

### 3. Run Tests

```bash
go test ./internal/...
```

### 4. Debug Tests

```bash
dlv test ./internal/user
```
