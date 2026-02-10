# Authentication Setup

## Overview
Simple JWT-based authentication system with register/login functionality.

## Features
- User registration with password hashing (bcrypt)
- User login with JWT token generation
- Protected routes with authentication middleware
- Session management via HTTP-only cookies
- JSON API endpoints for authentication

## Environment Variables
Add to your `.env` file:
```
JWT_SECRET=your-secret-key-here
```

## Routes

### Public Routes
- `GET /register` - Registration page
- `POST /register` - Create new user account
- `GET /login` - Login page
- `POST /login` - Authenticate user
- `POST /api/register` - JSON registration endpoint
- `POST /api/login` - JSON login endpoint

### Protected Routes (require authentication)
- `GET /` - Development tasks page
- `GET /notes` - Notes page
- `GET /health` - Health tracking page
- `GET /profile` - User profile page
- `GET /logout` - Logout and clear session

## How It Works

1. User registers with name, email, and password
2. Password is hashed using bcrypt
3. User data stored in MongoDB
4. On login, JWT token is generated and stored in HTTP-only cookie
5. Protected routes check for valid JWT token
6. Token expires after 7 days

## Database Collections
- `users` - User accounts with hashed passwords
- `tasks` - User tasks (linked by UserID)
- `notes` - User notes (linked by UserID)
- `sleep` - Sleep records (linked by UserID)
- `nutrition` - Nutrition records (linked by UserID)
- `activity` - Activity records (linked by UserID)

## Security Features
- Passwords hashed with bcrypt
- JWT tokens with expiration
- HTTP-only cookies (not accessible via JavaScript)
- User data isolated by UserID
- Protected routes redirect to login if not authenticated
