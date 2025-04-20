<<<<<<< HEAD
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
=======
  <h1>BlabberIt 🗨️</h1>
  <p>BlabberIt is a self-hosted, peer-to-peer chat application built in Go.<br>
  It features end-to-end encryption and stores messages locally when the recipient is offline.</p>

  <h2>Features</h2>
  <ul>
    <li>✅ User registration with public key</li>
    <li>✅ End-to-end encryption</li>
    <li>✅ Peer-to-peer messaging</li>
    <li>✅ Offline message queuing</li>
    <li>✅ PostgreSQL for persistent user data</li>
  </ul>

  <h2>Tech Stack</h2>
  <ul>
    <li>Go</li>
    <li>PostgreSQL</li>
    <li>GORM</li>
    <li>Docker (for DB)</li>
    <li>Delve (for debugging)</li>
    <li>Godotenv (for env config)</li>
  </ul>

  <h2>Getting Started</h2>

  <h3>1. Clone the repo</h3>
  <pre><code>git clone https://github.com/thechauhanabhay/blabberit.git
cd blabberit</code></pre>

  <h3>2. Setup <code>.env.test</code></h3>
  <pre><code>POSTGRES_DSN=host=localhost user=blabberuser password=blabberpass dbname=blabberit_test port=5432 sslmode=disable TimeZone=UTC</code></pre>

  <h3>3. Run Tests</h3>
  <pre><code>go test ./internal/...</code></pre>

  <h3>4. Debug Tests</h3>
  <pre><code>dlv test ./internal/user</code></pre>

</body>
</html>
>>>>>>> refs/remotes/origin/main
