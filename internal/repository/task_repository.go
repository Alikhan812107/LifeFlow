package repository

import "Assignment3/internal/models"

type TaskRepository struct {
	data map[int]models.Task
	last int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		data: make(map[int]models.Task),
	}
}

func (r *TaskRepository) Create(task models.Task) models.Task {
	r.last++
	task.ID = r.last
	r.data[task.ID] = task
	return task
}

func (r *TaskRepository) GetAll() []models.Task {
	tasks := make([]models.Task, 0)
	for _, t := range r.data {
		tasks = append(tasks, t)
	}
	return tasks
}

func (r *TaskRepository) GetByID(id int) (models.Task, bool) {
	task, ok := r.data[id]
	return task, ok
}

func (r *TaskRepository) Update(id int, task models.Task) (models.Task, bool) {
	if _, ok := r.data[id]; !ok {
		return models.Task{}, false
	}
	task.ID = id
	r.data[id] = task
	return task, true
}

func (r *TaskRepository) Delete(id int) bool {
	if _, ok := r.data[id]; !ok {
		return false
	}
	delete(r.data, id)
	return true
}
