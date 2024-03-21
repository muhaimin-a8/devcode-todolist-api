package usecase

import (
	"devcode-todolist-api/internal/domains"
	"devcode-todolist-api/internal/dtos"
	"log"
)

type todoUseCaseImpl struct {
	todoRepository     domains.TodoRepository
	activityRepository domains.ActivityRepository
}

func (t *todoUseCaseImpl) UpdateById(id string, req *dtos.TodoUpdateRequest) (isUpdated bool, res *dtos.TodoResponse, err error) {
	todoFromDB, err := t.todoRepository.GetById(id)
	if todoFromDB.Title == "" {
		return false, nil, nil
	}

	todo := domains.Todo{
		Id:       todoFromDB.Id,
		Title:    req.Title,
		Priority: req.Priority,
	}
	if req.Title != "" {
		todo.Title = req.Title
	}

	if req.Priority != "" {
		todo.Priority = req.Priority
	}

	updatedTodo, err := t.todoRepository.Update(todo)
	if err != nil {
		return false, nil, err
	}

	return true, &dtos.TodoResponse{
		Id:              updatedTodo.Id,
		ActivityGroupId: updatedTodo.ActivityGroupId,
		Title:           updatedTodo.Title,
		IsActive:        true,
		Priority:        updatedTodo.Priority,
		CreatedAt:       updatedTodo.CreatedAt,
		UpdatedAt:       updatedTodo.CreatedAt,
	}, nil
}

func (t *todoUseCaseImpl) DeleteById(id string) (isDeleted bool, err error) {
	return t.todoRepository.DeleteById(id)
}

func (t *todoUseCaseImpl) GetById(id string) (res *dtos.TodoResponse, err error) {
	todo, err := t.todoRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return &dtos.TodoResponse{
		Id:              todo.Id,
		ActivityGroupId: todo.ActivityGroupId,
		Title:           todo.Title,
		IsActive:        true,
		Priority:        todo.Priority,
		CreatedAt:       todo.CreatedAt,
		UpdatedAt:       todo.CreatedAt,
	}, nil
}

func (t *todoUseCaseImpl) GetAllByActivityId(activityId string) (isActivityExist bool, res *[]dtos.TodoResponse, err error) {
	// check if activity group is exist
	activity, err := t.activityRepository.GetById(activityId)
	if err != nil {
		return false, nil, err
	}

	if activity.Title == "" {
		return false, nil, nil
	}

	todos, err := t.todoRepository.GetAllByActivityId(activityId)
	if err != nil {
		return true, nil, err
	}

	var listResponse []dtos.TodoResponse
	for i := 0; i < len(todos); i++ {
		res := dtos.TodoResponse{
			Id:              todos[i].Id,
			ActivityGroupId: todos[i].ActivityGroupId,
			Title:           todos[i].Title,
			IsActive:        true,
			Priority:        todos[i].Priority,
			CreatedAt:       todos[i].CreatedAt,
			UpdatedAt:       todos[i].CreatedAt,
		}
		listResponse = append(listResponse, res)
	}

	return true, &listResponse, err
}

func (t *todoUseCaseImpl) CreateNew(req *dtos.TodoCreateRequest) (bool, *dtos.TodoResponse, error) {
	// check if activity group is exist
	activity, err := t.activityRepository.GetById(req.ActivityGroupId)
	if err != nil {
		return false, nil, err
	}

	if activity.Title == "" {
		return false, nil, nil
	}

	todo := domains.Todo{
		ActivityGroupId: req.ActivityGroupId,
		Title:           req.Title,
		Priority:        req.Priority,
	}

	todoFromDB, err := t.todoRepository.Save(todo)
	if err != nil {
		return true, nil, err
	}
	log.Println(todoFromDB)

	return true, &dtos.TodoResponse{
		Id:              todoFromDB.Id,
		ActivityGroupId: todoFromDB.ActivityGroupId,
		Title:           todoFromDB.Title,
		IsActive:        true,
		Priority:        todoFromDB.Priority,
		CreatedAt:       todoFromDB.CreatedAt,
		UpdatedAt:       todoFromDB.CreatedAt,
	}, nil
}

func NewTodoUseCase(
	todoRepository domains.TodoRepository,
	activityRepository domains.ActivityRepository,
) domains.TodoUseCase {
	return &todoUseCaseImpl{
		todoRepository:     todoRepository,
		activityRepository: activityRepository,
	}
}
