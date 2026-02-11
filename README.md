# LifeFlow - Personal Productivity & Health Tracker

Simple web app for managing tasks, notes, and health tracking built with Go and MongoDB.

## Features

- Create and organize tasks with folders
- Mark tasks as complete/incomplete
- Create simple notes (title + description)
- Track sleep patterns (bedtime and wake time)
- Track nutrition (calories, water intake, healthy eating)
- Log daily activities
- View user profile with statistics
- Upload custom avatar image
- Full CRUD operations for all entities
- Interactive charts using Chart.js

## Quick Start

1. Start MongoDB
2. Copy `.env.example` to `.env` and set your `MONGO_URI`
3. Run: `go run cmd/server/main.go`
4. Open: http://localhost:8080

## Pages

- **Development** (`/`) - Task management with folders and inline editing
- **Notes** (`/notes`) - Simple note-taking without status tracking
- **Health** (`/health`) - Track sleep, nutrition, and activities with charts
- **Profile** (`/profile`) - User info, statistics, and avatar upload

## Architecture

3-layer architecture pattern:

```
Handler -> Service -> Repository -> MongoDB
```

- **Handler**: HTTP request handling, form parsing, template rendering
- **Service**: Business logic layer (currently thin, delegates to repository)
- **Repository**: Database operations using MongoDB driver
- **Models**: Data structures with JSON and BSON tags

## API Endpoints

### Tasks
- `GET /` - Main page with tasks
- `POST /development/html` - Create task
- `POST /development/update` - Update task
- `GET /development/toggle?id=<id>` - Toggle task status
- `GET /development/delete?id=<id>` - Delete task

### Notes
- `GET /notes` - Notes page
- `POST /notes/html` - Create note
- `POST /notes/update` - Update note
- `GET /notes/delete?id=<id>` - Delete note

### Health
- `GET /health` - Health tracking page
- `POST /health/sleep` - Add sleep record
- `POST /health/nutrition` - Add nutrition record
- `POST /health/activity` - Add activity record

### Profile
- `GET /profile` - User profile page
- `POST /profile/avatar` - Upload avatar image (max 10MB)

## Data Models

```go
type Task struct {
    ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title  string             `json:"title" bson:"title"`
    Body   string             `json:"body" bson:"body"`
    Done   bool               `json:"done" bson:"done"`
    Folder string             `json:"folder" bson:"folder"`
    UserID string             `json:"user_id" bson:"user_id"`
}

type Note struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title       string             `json:"title" bson:"title"`
    Description string             `json:"description" bson:"description"`
    UserID      string             `json:"user_id" bson:"user_id"`
}

type Sleep struct {
    ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    WokeUp    time.Time          `json:"woke_up" bson:"woke_up"`
    Slept     time.Time          `json:"slept" bson:"slept"`
    UserID    string             `json:"user_id" bson:"user_id"`
    Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}

type Nutrition struct {
    ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Calories  int                `json:"calories" bson:"calories"`
    Water     float64            `json:"water" bson:"water"`
    Healthy   bool               `json:"healthy" bson:"healthy"`
    UserID    string             `json:"user_id" bson:"user_id"`
    Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}

type Activity struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Description string             `json:"description" bson:"description"`
    UserID      string             `json:"user_id" bson:"user_id"`
    Timestamp   time.Time          `json:"timestamp" bson:"timestamp"`
}

type User struct {
    ID       string `json:"id" bson:"_id,omitempty"`
    Name     string `json:"name" bson:"name"`
    Email    string `json:"email" bson:"email"`
    TasksNum int    `json:"tasks_num" bson:"tasks_num"`
    NotesNum int    `json:"notes_num" bson:"notes_num"`
    Avatar   string `json:"avatar" bson:"avatar"`
}
```

## MongoDB Collections

- `tasks` - User tasks with folder organization
- `notes` - Simple notes
- `sleep` - Sleep tracking records
- `nutrition` - Nutrition tracking records
- `activity` - Activity logs
- `users` - User profiles with avatar (stored as base64)

## Tech Stack

- Go 1.x
- MongoDB (with official Go driver)
- HTML templates
- Chart.js for data visualization
- Pure CSS (no frameworks)

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── internal/
│   ├── app/
│   │   ├── router.go            # Route registration
│   │   └── server.go            # HTTP server
│   ├── handlers/                # HTTP handlers
│   ├── service/                 # Business logic
│   ├── repository/              # Database operations
│   └── models/                  # Data structures
└── templates/                   # HTML templates
```

## How Avatar Upload Works

1. User selects image file on profile page
2. Form submits to `/profile/avatar` with multipart/form-data
3. Handler reads file bytes (max 10MB)
4. Converts image to base64 string
5. Saves to MongoDB users collection
6. On profile load, retrieves base64 string
7. Displays as `<img src="data:image/jpeg;base64,...">`

## Environment Variables

```
MONGO_URI=mongodb://localhost:27017
```
