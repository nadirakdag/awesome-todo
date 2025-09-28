package todo

import (
	"fmt"
	"time"
)

type Service interface {
	GetList() ([]Task, error)
	GetById(id string) (Task, error)
	Create(description string) (Task, error)
	Complete(id string) (Task, error)
	Delete(id string) (Task, error)
	Edit(id string, description string) (Task, error)
}

type taskService struct{}

func NewService() Service {
	return &taskService{}
}

func (ts *taskService) GetList() ([]Task, error) {
	tasks := make([]Task, 10)

	for i := 0; i < len(tasks); i++ {
		tasks[i] = Task{
			ID:          fmt.Sprintf("fdsfs%d", i),
			Description: fmt.Sprintf("Test Task Description %d", i),
			Status:      false,
			CreatedAt:   time.Now().AddDate(0, 0, i),
		}
	}

	return tasks, nil
}
func (ts *taskService) GetById(id string) (Task, error) {
	return Task{}, fmt.Errorf("not implemented")
}
func (ts *taskService) Create(description string) (Task, error) {
	return Task{}, fmt.Errorf("not implemented")
}
func (ts *taskService) Complete(id string) (Task, error) {
	return Task{}, fmt.Errorf("not implemented")
}
func (ts *taskService) Delete(id string) (Task, error) {
	return Task{}, fmt.Errorf("not implemented")
}
func (ts *taskService) Edit(id string, description string) (Task, error) {
	return Task{}, fmt.Errorf("not implemented")
}
