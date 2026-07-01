# Go Fiber v3 API - Learning Management System

## Getting Started

### 1. Requirements
*   **Go** version 1.23 or higher

### 2. Configuration
Copy the configuration template to create your local environment variables:
```bash
cp .env.example .env
```
Open the `.env` file and set the required variables:
```env
APP_PORT=8080

JWT_SECRET_KEY=your_very_secure_jwt_secret_key
JWT_EXPIRATION_HOURS=24
JWT_REFRESH_EXPIRATION_DAYS=14

REDIS_URL=""

# Neon PostgreSQL connection string (or local postgres database URL)
DATABASE_URL="postgresql://username:password@host:port/database?sslmode=require"

# Optional SMTP Email Configuration
SMTP_HOST=smtp.mailtrap.io
SMTP_PORT=2525
SMTP_USER=your_smtp_user
SMTP_PASSWORD=your_smtp_password
SMTPSender=no-reply@example.com
```
