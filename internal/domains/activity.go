package domains

import (
	"devcode-todolist-api/internal/dtos"
)

type Activity struct {
	Id        int    `db:"activity_id"`
	Title     string `db:"title"`
	Email     string `db:"email"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	DeletedAt string `db:"deleted_at"`
}

type ActivityUseCase interface {
	CreateNew(req *dtos.ActivityCreateRequest) (res *dtos.ActivityResponse, err error)
	GetAll() (res *[]dtos.ActivityResponse, err error)
	GetById(id int) (res *dtos.ActivityResponse, err error)
	DeleteById(id int) (isDeleted bool, err error)
	UpdateById(id int, req *dtos.ActivityUpdateRequest) (isUpdated bool, res *dtos.ActivityResponse, err error)
}

type ActivityRepository interface {
	Save(activity Activity) (*Activity, error)
	GetAll() ([]Activity, error)
	GetById(id int) (*Activity, error)
	DeleteById(id int) (isDeleted bool, err error)
	Update(activity Activity) (*Activity, error)
}
