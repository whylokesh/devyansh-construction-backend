# 🏗️Devyansh Construction Workforce & Site Management System Backend

A complete system to manage construction sites, workers, attendance, billing, payouts, advances, and site summaries for labour contractors and manpower agencies.

This project solves the real-world problem of daily worker movement between sites, attendance-driven billing, and flexible monthly payouts with partial advance deductions.

## Features

### 👥 User Management
- **Roles**: Supports `admin` and `accountant` roles.
- **Authentication**: Secure password handling and user sessions.

### 👷 Worker Management
- **Profiles**: Manage worker details, skills, and contact info.
- **Rates**: Track individual bill rates and payout rates.
- **Status**: Monitor active/inactive status of workers.

### 🏗️ Site Management
- **Site Tracking**: Manage multiple construction sites with location and client details.
- **Status**: Track site status (`active`, `completed`).
- **Documents**: Store and reference site-related documents.

### 📅 Attendance & Tracking
- **Daily Logs**: Record worker attendance (present, absent, half-day).
- **Site-Specific**: Track which worker is at which site on a given day.

### 💰 Financials
- **Advances**: Manage advance payments given to workers.
- **Payouts**: Calculate and track worker payouts for specific periods.
- **Bills**: Generate bills for clients based on site work.
- **Status Tracking**: Track status of bills and payouts (`draft`, `saved`, `paid`, `partial`).

## Tech Stack
- **Language**: [Go](https://go.dev/) (v1.24)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **Router**: [Chi](https://github.com/go-chi/chi)
- **Migrations**: [Golang-Migrate](https://github.com/golang-migrate/migrate)

## Project Structure
```
.
├── cmd/
│   └── server/         # Application entry point
├── internal/
│   ├── config/         # Configuration loading
│   ├── db/             # Database connection and migrations
│   ├── handler/        # HTTP request handlers
│   ├── middleware/     # HTTP middlewares
│   ├── models/         # Data models
│   ├── repository/     # Data access layer
│   ├── routes/         # Route definitions
│   └── service/        # Business logic
├── references/         # Project references and docs
├── .env                # Environment variables
├── Makefile            # Build and utility commands
└── go.mod              # Go module definition
```

## Getting Started

### Prerequisites
- Go 1.24+
- PostgreSQL
- Make (optional, for running Makefile commands)

### Configuration
1. Create a `.env` file in the root directory (or use system env vars).
2. Configure the following variables:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=devyansh_construction
   ```

### Database Setup
Run the migrations to set up the database schema:
```bash
make migrate-up
```

### Running the Application
Start the server:
```bash
make run
```
The server will start on port `8080` (default).

## API Endpoints
(Documentation for specific API endpoints can be added here as they are developed)

## License
Private - Devyansh Construction
