package domains

import "devcode-todolist-api/internal/dtos"

type Todo struct {
	Id              int    `db:"todo_id"`
	ActivityGroupId string `db:"activity_group_id"`
	Title           string `db:"title"`
	Priority        string `db:"priority"`
	CreatedAt       string `db:"created_at"`
	//UpdatedAt string `db:"updated_at"`
	//DeletedAt string `db:"deleted_at"`
}

type TodoUseCase interface {
	CreateNew(req *dtos.TodoCreateRequest) (isActivityExist bool, res *dtos.TodoResponse, err error)
	GetAllByActivityId(activityId string) (isActivityExist bool, res *[]dtos.TodoResponse, err error)
	GetById(id string) (res *dtos.TodoResponse, err error)
	DeleteById(id string) (isDeleted bool, err error)
	UpdateById(id string, req *dtos.TodoUpdateRequest) (isUpdated bool, res *dtos.TodoResponse, err error)
}

type TodoRepository interface {
	Save(todo Todo) (*Todo, error)
	GetAllByActivityId(activityId string) ([]Todo, error)
	GetById(id string) (*Todo, error)
	DeleteById(id string) (isDeleted bool, err error)
	Update(todo Todo) (*Todo, error)
}
