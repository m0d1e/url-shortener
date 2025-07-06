# 🔗 URL Shortener

A fast, minimalistic, and extensible URL shortening service built with Go and PostgreSQL.

---

## ✨ Features

* 🔧 Shorten long URLs into simple aliases
* 🥉 Support for custom or random aliases
* ⚙️ Clean RESTful API
* 📔 PostgreSQL-backed persistent storage
* 🔐 Basic authentication for protected operations (optional)
* 📦 Modular and clean architecture
* 📊 Structured logging with `slog`
* ✅ Input validation
* ⚡ Fast and lightweight with `chi`

---

## 🚀 Tech Stack

| Tool                                                    | Purpose            |
| ------------------------------------------------------- | ------------------ |
| [Go](https://golang.org/)                               | Core language      |
| [PostgreSQL](https://www.postgresql.org/)               | Persistent storage |
| [chi](https://github.com/go-chi/chi)                    | HTTP routing       |
| [pgx](https://github.com/jackc/pgx)                     | PostgreSQL driver  |
| [validator](https://github.com/go-playground/validator) | Request validation |
| [slog](https://pkg.go.dev/log/slog)                     | Structured logging |
| [cleanenv](https://github.com/ilyakaznacheev/cleanenv)  | Config loader      |
| [godotenv](https://github.com/joho/godotenv)            | `.env` file loader |

---

## 🧱 Project Structure

```
.
├── cmd/
│   └── url-shortener/      # Application entry point
├── internal/
│   ├── api/                # API server and routing
│   ├── config/             # Configuration loader
│   ├── http/
│   │   ├── dto/            # Request/response DTOs
│   │   └── handlers/       # HTTP handlers
│   ├── middleware/         # Middleware
│   ├── slErr/              # Logging helpers
│   └── storage/            # Database access
├── service/
│   ├── generateAlias/      # Alias generation logic
│   ├── repo.go             # Repository interface
│   └── url_service.go      # Business logic
├── config.example.yaml     # Example config
├── .env.example            # Example environment variables
├── go.mod / go.sum         # Go modules
├── LICENSE                 # License file
├── README.md
└── .gitignore
```

---

## ⚙️ Getting Started

### 🥉 Prerequisites

* Go 1.20+
* PostgreSQL
* Git

### 🛠 Installation

```bash
# Clone the repo
git clone https://github.com/m0d1e/url-shortener.git
cd url-shortener

# Copy config files
cp .env.example .env
cp config.example.yaml config.yaml

# Run the app
go run ./cmd/url-shortener
```

---

## 📡 API Endpoints

| Method   | Endpoint         | Description              |
| -------- | ---------------- | ------------------------ |
| POST     | `/url/`          | Shorten a URL (requires Basic Auth)         |
| GET      | `/{alias}`       | Redirect to the original URL |
| DELETE   | `/url/{alias}`   | Delete a shortened URL (requires Basic Auth) |

### 🔐 Optional Authentication

Basic authentication headers can be enabled for protected routes (configurable in `.env`).

---

## 🤝 Contributing

Contributions welcome! Open issues, create pull requests, or suggest improvements.

---

## 🧪 License

This project is licensed under the MIT License. See `LICENSE` file for details.