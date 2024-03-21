package usecase

import (
	"devcode-todolist-api/internal/domains"
	"devcode-todolist-api/internal/dtos"
	"time"
)

type activityUseCaseImpl struct {
	repository domains.ActivityRepository
}

func (a *activityUseCaseImpl) UpdateById(id string, req *dtos.ActivityUpdateRequest) (isUpdated bool, res *dtos.ActivityResponse, err error) {
	activityFromDB, err := a.repository.GetById(id)
	if activityFromDB.Title == "" {
		return false, nil, nil
	}

	activity := domains.Activity{
		Id:    activityFromDB.Id,
		Title: activityFromDB.Title,
		Email: activityFromDB.Email,
	}
	if req.Title != "" {
		activity.Title = req.Title
	}

	if req.Email != "" {
		activity.Email = req.Email
	}

	updatedActivity, err := a.repository.Update(activity)
	if err != nil {
		return false, nil, err
	}

	return true, &dtos.ActivityResponse{
		Id:        updatedActivity.Id,
		Title:     updatedActivity.Title,
		Email:     updatedActivity.Email,
		CreatedAt: updatedActivity.CreatedAt,
		UpdatedAt: time.Now().String(),
	}, nil
}

func (a *activityUseCaseImpl) DeleteById(id string) (isDeleted bool, err error) {
	return a.repository.DeleteById(id)
}

func (a *activityUseCaseImpl) GetById(id string) (res *dtos.ActivityResponse, err error) {
	activityFromDB, err := a.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return &dtos.ActivityResponse{
		Id:        activityFromDB.Id,
		Title:     activityFromDB.Title,
		Email:     activityFromDB.Email,
		CreatedAt: activityFromDB.CreatedAt,
		UpdatedAt: activityFromDB.CreatedAt,
	}, nil
}

func (a *activityUseCaseImpl) GetAll() (res *[]dtos.ActivityResponse, err error) {
	listActivity, err := a.repository.GetAll()
	if err != nil {
		return nil, err
	}

	var listResponse []dtos.ActivityResponse
	for i := 0; i < len(listActivity); i++ {
		res := dtos.ActivityResponse{
			Id:        listActivity[i].Id,
			Title:     listActivity[i].Title,
			Email:     listActivity[i].Email,
			CreatedAt: listActivity[i].CreatedAt,
			UpdatedAt: listActivity[i].CreatedAt,
		}
		listResponse = append(listResponse, res)
	}

	return &listResponse, err
}

func (a *activityUseCaseImpl) CreateNew(req *dtos.ActivityCreateRequest) (res *dtos.ActivityResponse, err error) {
	activity := domains.Activity{
		Title: req.Title,
		Email: req.Email,
	}

	activityFromDB, err := a.repository.Save(activity)
	if err != nil {
		return nil, err
	}

	return &dtos.ActivityResponse{
		Id:        activityFromDB.Id,
		Title:     activityFromDB.Title,
		Email:     activityFromDB.Email,
		CreatedAt: activityFromDB.CreatedAt,
		UpdatedAt: activityFromDB.CreatedAt,
	}, nil
}

func NewActivityUseCase(repository domains.ActivityRepository) domains.ActivityUseCase {
	return &activityUseCaseImpl{repository: repository}
}
