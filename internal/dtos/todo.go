package dtos

type TodoCreateRequest struct {
	ActivityGroupId string `json:"activity_group_id" validate:"required"`
	Title           string `json:"title" validate:"required,min=2,max=255"`
	Priority        string `json:"priority" validate:"min=2,max=20"`
}

type TodoResponse struct {
	Id              int    `json:"id,omitempty"`
	ActivityGroupId string `json:"activity_group_id,omitempty"`
	Title           string `json:"title,omitempty"`
	IsActive        bool   `json:"is_active,omitempty"`
	Priority        string `json:"priority,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
}

type TodoUpdateRequest struct {
	Title    string `json:"title" validate:"omitempty,max=255"`
	Priority string `json:"priority" validate:"omitempty,max=20"`
}
