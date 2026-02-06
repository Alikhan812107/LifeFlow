# LifeFlow - Task & Notes App

Простое приложение для тасков и заметок с Go и MongoDB.

## Что умеет

- создавать таски с папками
- отмечать таски как выполненные
- создавать простые заметки (без тасков)
- трекать сон (когда лег спать и проснулся)
- трекать питание (калории, вода, здоровое питание)
- записывать активности
- смотреть профиль пользователя
- все операции CRUD
- графики через Chart.js

## Как запустить

1. запустить mongodb
2. скопировать `.env.example` в `.env` 
3. запустить: `go run cmd/server/main.go`
4. открыть: http://localhost:8080

## Страницы

- **Tasks** (`/`) - таски с папками и inline редактированием
- **Notes** (`/notes`) - простые заметки без статуса
- **Sleep** (`/sleep`) - трекинг сна с графиком длительности
- **Nutrition** (`/nutrition`) - трекинг питания с графиками калорий, воды и здорового питания
- **Activity** (`/activity`) - запись активностей
- **Profile** (`/profile`) - инфо о пользователе и статистика

## Что добавлено

- папки для тасков (можно сортировать)
- отдельная страница заметок (title + description)
- страница профиля с статистикой
- трекинг сна с временем засыпания и пробуждения
- трекинг питания (калории, вода, здоровое питание да/нет)
- запись активностей с временными метками
- графики через Chart.js для визуализации данных
- все данные хранятся в MongoDB как time series
- навигация между страницами
- все CRUD операции на каждой странице
- код написан как студент

## API

**Tasks:**
- `GET /` - главная страница с тасками
- `POST /tasks/html` - создать таск
- `POST /tasks/update` - обновить таск
- `GET /tasks/toggle?id=<id>` - переключить статус
- `GET /tasks/delete?id=<id>` - удалить таск

**Notes:**
- `GET /notes` - страница заметок
- `POST /notes/html` - создать заметку
- `POST /notes/update` - обновить заметку
- `GET /notes/delete?id=<id>` - удалить заметку

**Sleep:**
- `GET /sleep` - страница трекинга сна
- `POST /sleep/html` - добавить запись сна
- `GET /sleep/json` - получить все записи в JSON

**Nutrition:**
- `GET /nutrition` - страница трекинга питания
- `POST /nutrition/html` - добавить запись питания
- `GET /nutrition/json` - получить все записи в JSON

**Activity:**
- `GET /activity` - страница активностей
- `POST /activity/html` - добавить активность
- `GET /activity/json` - получить все активности в JSON

**Profile:**
- `GET /profile` - профиль пользователя

## Структуры

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
}
```