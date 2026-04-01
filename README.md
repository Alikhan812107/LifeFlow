# LifeFlow - Personal Life Management System

A comprehensive web application built with Go that helps you manage your daily life including tasks, notes, health tracking (sleep, nutrition, activities), and more. Perfect for students learning Go web development!

## 📋 Table of Contents
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Database Schema](#database-schema)
- [Architecture](#architecture)
- [Learning Resources](#learning-resources)

## ✨ Features

### User Management
- User registration and authentication
- Session-based login system
- User profiles with role-based access (Free/Premium)
- Password hashing with bcrypt

### Notes Management
- Create, read, update, and delete notes
- Free users: Limited to 10 notes
- Premium users: Unlimited notes
- Each note has a title and description

### Task Management
- Create and manage tasks
- Track task completion status
- Organize your to-do list

### Health Tracking
- **Sleep Tracking**: Monitor your sleep patterns
- **Nutrition Tracking**: Log your meals and dietary information
- **Activity Tracking**: Record physical activities and exercises

### Subscription System
- Free tier with limited features
- Premium tier with unlimited access
- Middleware-based subscription enforcement

## 🛠 Tech Stack

- **Language**: Go 1.25
- **Database**: MongoDB (Cloud - MongoDB Atlas)
- **Web Framework**: Native Go `net/http`
- **Template Engine**: Go `html/template`
- **Authentication**: Session-based with bcrypt password hashing
- **Environment Management**: godotenv

### Key Dependencies
```go
github.com/joho/godotenv v1.5.1          // Environment variable management
go.mongodb.org/mongo-driver v1.17.7      // MongoDB driver
golang.org/x/crypto v0.48.0              // Password hashing
```

## 📁 Project Structure

```
Assignment3/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── app/
│   │   ├── router.go              # Route definitions
│   │   └── server.go              # HTTP server setup
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── handlers/                  # HTTP request handlers
│   │   ├── auth_handler.go        # Login/Register handlers
│   │   ├── note_handler.go        # Note CRUD handlers
│   │   ├── user_handler.go        # User profile handlers
│   │   ├── health_handler.go      # Health tracking handlers
│   │   └── development_handler.go # Development utilities
│   ├── middleware/                # HTTP middleware
│   │   ├── auth_middleware.go     # Authentication checks
│   │   ├── session_middleware.go  # Session management
│   │   ├── subscription_middleware.go # Subscription checks
│   │   └── logging_middleware.go  # Request logging
│   ├── models/                    # Data models
│   │   ├── user.go
│   │   ├── note.go
│   │   ├── task.go
│   │   ├── sleep.go
│   │   ├── nutrition.go
│   │   └── activity.go
│   ├── repository/                # Database layer
│   │   ├── mongo_client.go        # MongoDB connection
│   │   └── *_repository.go        # Repository implementations
│   └── service/                   # Business logic layer
│       └── *_service.go           # Service implementations
├── templates/                     # HTML templates
│   ├── login.html
│   ├── register.html
│   ├── notes.html
│   ├── profile.html
│   ├── health.html
│   └── style.css
├── .env                          # Environment variables (not in git)
├── .gitignore
├── go.mod                        # Go module definition
└── README.md
```

## 📚 Architecture

This project follows a **clean architecture** pattern with clear separation of concerns:

1. **Handlers Layer** (`internal/handlers/`)
   - Handles HTTP requests and responses
   - Validates input data
   - Calls service layer for business logic

2. **Service Layer** (`internal/service/`)
   - Contains business logic
   - Orchestrates data flow between handlers and repositories
   - Enforces business rules

3. **Repository Layer** (`internal/repository/`)
   - Handles database operations
   - Abstracts database implementation details
   - Uses interfaces for flexibility

4. **Models Layer** (`internal/models/`)
   - Defines data structures
   - Represents database entities

5. **Middleware Layer** (`internal/middleware/`)
   - Authentication checks
   - Session management
   - Logging and monitoring

## 🔧 Prerequisites

Before you begin, ensure you have the following installed:

- **Go**: Version 1.25 or higher
  - Download from [golang.org](https://golang.org/dl/)
  - Verify installation: `go version`

- **MongoDB Atlas Account** (Free tier works!)
  - Sign up at [mongodb.com/cloud/atlas](https://www.mongodb.com/cloud/atlas)
  - Create a cluster and get your connection string

- **Git**: For version control
  - Download from [git-scm.com](https://git-scm.com/)

## 🚀 Installation

### Step 1: Clone the Repository
```bash
git clone <your-repository-url>
cd Assignment3
```

### Step 2: Install Dependencies
```bash
go mod download
```

This will download all required packages listed in `go.mod`.

### Step 3: Verify Installation
```bash
go mod verify
```

## ⚙️ Configuration

### Step 1: Create Environment File
Create a `.env` file in the project root:

```bash
# On Windows (CMD)
type nul > .env

# On Windows (PowerShell)
New-Item .env -ItemType File
```

### Step 2: Add Environment Variables
Open `.env` and add the following:

```env
MONGO_URI=mongodb+srv://<username>:<password>@<cluster>.mongodb.net/?appName=<AppName>
PORT=8080
```

**Important**: Replace the placeholders:
- `<username>`: Your MongoDB username
- `<password>`: Your MongoDB password
- `<cluster>`: Your cluster address
- `<AppName>`: Your application name

### Step 3: MongoDB Setup
1. Log in to [MongoDB Atlas](https://cloud.mongodb.com/)
2. Create a new cluster (free tier is fine)
3. Create a database user with read/write permissions
4. Whitelist your IP address (or use `0.0.0.0/0` for development)
5. Get your connection string from "Connect" → "Connect your application"
6. The database `lifeflow` and collections will be created automatically

## 🏃 Running the Application

### Development Mode

#### Option 1: Using go run
```bash
go run cmd/server/main.go
```

#### Option 2: Build and Run
```bash
# Build the executable
go build -o lifeflow.exe cmd/server/main.go

# Run the executable
.\lifeflow.exe
```

### Expected Output
```
server starting on :8080
```

### Access the Application
Open your browser and navigate to:
```
http://localhost:8080
```

## 🌐 API Endpoints

### Authentication
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/login` | Login page | No |
| POST | `/login` | Login user | No |
| GET | `/register` | Registration page | No |
| POST | `/register` | Register new user | No |
| GET | `/logout` | Logout user | Yes |

### Notes
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/notes` | View all notes (HTML) | Yes |
| POST | `/notes/create` | Create new note | Yes |
| POST | `/notes/update` | Update existing note | Yes |
| GET | `/notes/delete` | Delete note | Yes |

### User Profile
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/profile` | View user profile | Yes |
| POST | `/profile/update` | Update profile | Yes |

### Health Tracking
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/health` | Health dashboard | Yes |
| POST | `/health/sleep` | Log sleep data | Yes |
| POST | `/health/nutrition` | Log nutrition data | Yes |
| POST | `/health/activity` | Log activity data | Yes |

### Development (Debug endpoints)
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/dev` | Development dashboard | No |
| GET | `/dev/users` | List all users | No |
| POST | `/dev/clear-sessions` | Clear all sessions | No |

## 🗄️ Database Schema

### Database Name: `lifeflow`

### Collections:

#### users
```json
{
  "_id": "ObjectId",
  "username": "string",
  "email": "string",
  "password": "string (hashed)",
  "role": "string (free/premium)",
  "createdAt": "timestamp"
}
```

#### notes
```json
{
  "_id": "ObjectId",
  "title": "string",
  "description": "string",
  "userId": "ObjectId",
  "createdAt": "timestamp"
}
```

#### tasks
```json
{
  "_id": "ObjectId",
  "title": "string",
  "description": "string",
  "completed": "boolean",
  "userId": "ObjectId",
  "createdAt": "timestamp"
}
```

#### sleep
```json
{
  "_id": "ObjectId",
  "userId": "ObjectId",
  "date": "timestamp",
  "hours": "number",
  "quality": "string"
}
```

#### nutrition
```json
{
  "_id": "ObjectId",
  "userId": "ObjectId",
  "date": "timestamp",
  "meal": "string",
  "calories": "number"
}
```

#### activity
```json
{
  "_id": "ObjectId",
  "userId": "ObjectId",
  "date": "timestamp",
  "type": "string",
  "duration": "number"
}
```

