package domains

import "devcode-todolist-api/internal/dtos"

type Todo struct {
	Id              int    `db:"todo_id"`
	ActivityGroupId int    `db:"activity_group_id"`
	Title           string `db:"title"`
	Priority        string `db:"priority"`
	CreatedAt       string `db:"created_at"`
	//UpdatedAt string `db:"updated_at"`
	//DeletedAt string `db:"deleted_at"`
}

type TodoUseCase interface {
	CreateNew(req *dtos.TodoCreateRequest) (isActivityExist bool, res *dtos.TodoResponse, err error)
	GetAllByActivityId(activityId int) (isActivityExist bool, res *[]dtos.TodoResponse, err error)
	GetById(id int) (res *dtos.TodoResponse, err error)
	DeleteById(id int) (isDeleted bool, err error)
	UpdateById(id int, req *dtos.TodoUpdateRequest) (isUpdated bool, res *dtos.TodoResponse, err error)
}

type TodoRepository interface {
	Save(todo Todo) (*Todo, error)
	GetAllByActivityId(activityId int) ([]Todo, error)
	GetById(id int) (*Todo, error)
	DeleteById(id int) (isDeleted bool, err error)
	Update(todo Todo) (*Todo, error)
}
