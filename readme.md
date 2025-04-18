<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>BlabberIt</title>
  <style>
    body { font-family: sans-serif; line-height: 1.6; max-width: 800px; margin: auto; padding: 2rem; }
    h1, h2, h3 { color: #333; }
    code { background: #f4f4f4; padding: 0.2em 0.4em; border-radius: 4px; }
    pre { background: #f4f4f4; padding: 1em; border-radius: 6px; overflow-x: auto; }
  </style>
</head>
<body>

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

  <h2>Directory Structure</h2>
  <pre><code>internal/
├── user/         → User handlers, models, tests
├── testutils/    → DB setup for tests
cmd/
└── server/       → App entrypoint (coming soon)</code></pre>

  <h2>License</h2>
  <p>MIT</p>

</body>
</html>
