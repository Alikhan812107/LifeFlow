# Authentication Fix - User Data Isolation

## Problem Found
After registration and login, users could see ALL data from ALL users instead of just their own data.

## Root Cause
The repository layer was fetching all records without filtering by user ID:
```go
cursor, err := r.collection.Find(context.Background(), bson.M{})
```

## Solution Applied
Updated all repositories, services, and handlers to filter data by authenticated user ID:

### Changes Made:

1. **Repository Interfaces** - Added userID parameter to GetAll methods
   - TaskRepository
   - NoteRepository
   - SleepRepository
   - NutritionRepository
   - ActivityRepository

2. **Repository Implementations** - Filter by user_id in MongoDB queries
   ```go
   cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": userID})
   ```

3. **Services** - Pass userID to repository methods
   - TaskService.GetAll(userID)
   - NoteService.GetAll(userID)
   - SleepService.GetAll(userID)
   - NutritionService.GetAll(userID)
   - ActivityService.GetAll(userID)

4. **Handlers** - Extract userID from request context and pass to services
   ```go
   userID := middleware.GetUserID(r)
   tasks, err := h.service.GetAll(userID)
   ```

## Result
Now each user only sees their own:
- Development tasks
- Notes
- Sleep records
- Nutrition records
- Activity records

## How It Works Now

1. User logs in
2. JWT token with user ID is stored in cookie
3. Middleware extracts user ID from token
4. Handlers get user ID from request context
5. Services and repositories filter all queries by user ID
6. User only sees their own data

## Testing
1. Register a new account
2. Login with your credentials
3. Create some tasks/notes
4. Logout and register another account
5. Login with second account
6. You should see empty pages (no data from first user)
7. Create different tasks/notes
8. Each user sees only their own data
