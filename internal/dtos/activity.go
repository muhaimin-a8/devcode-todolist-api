package dtos

type ActivityCreateRequest struct {
	Title string `json:"title" validate:"required,min=2,max=255"`
	Email string `json:"email" validate:"required,email"`
}

type ActivityResponse struct {
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Email     string `json:"email,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

type ActivityUpdateRequest struct {
	Title string `json:"title" validate:"omitempty,max=255"`
	Email string `json:"email" validate:"omitempty,email"`
}
