# LifeFlow - Task & Notes App

Простое приложение для тасков и заметок с Go и MongoDB.

## Что умеет

- создавать таски с папками
- отмечать таски как выполненные
- создавать простые заметки (без тасков)
- смотреть профиль пользователя
- все операции CRUD

## Как запустить

1. запустить mongodb
2. скопировать `.env.example` в `.env` 
3. запустить: `go run cmd/server/main.go`
4. открыть: http://localhost:8080

## Страницы

- **Tasks** (`/`) - таски с папками и inline редактированием
- **Notes** (`/notes`) - простые заметки без статуса
- **Profile** (`/profile`) - инфо о пользователе и статистика

## Что добавлено

- папки для тасков (можно сортировать)
- отдельная страница заметок (title + description)
- страница профиля с статистикой
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

type User struct {
    ID       string `json:"id" bson:"_id,omitempty"`
    Name     string `json:"name" bson:"name"`
    Email    string `json:"email" bson:"email"`
    TasksNum int    `json:"tasks_num" bson:"tasks_num"`
    NotesNum int    `json:"notes_num" bson:"notes_num"`
}
```