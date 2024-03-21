package api

import (
	"devcode-todolist-api/internal/domains"
	"devcode-todolist-api/internal/dtos"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterActivityController(router fiber.Router, validate *validator.Validate, usecase domains.ActivityUseCase) {
	// create new activity group
	router.Post("/activity-groups", func(ctx *fiber.Ctx) error {
		res := ctx.Response()
		res.Header.Set("Content-Type", "application/json")

		var req *dtos.ActivityCreateRequest
		err := ctx.BodyParser(&req)
		if err != nil {
			// failed to parse request body
			data, _ := json.Marshal(dtos.Response{
				Status:  "error",
				Message: "can not parse request body",
			})

			res.SetBody(data)
			res.SetStatusCode(400)

			return nil
		}

		err = validate.Struct(req)
		if err != nil {
			// failed to validate request body
			data, _ := json.Marshal(dtos.Response{
				Status:  "failed",
				Message: err.Error(),
			})

			res.SetBody(data)
			res.SetStatusCode(401)

			return nil
		}

		activity, err := usecase.CreateNew(req)
		data, _ := json.Marshal(dtos.Response{
			Status:  "success",
			Message: "success to create new activity group",
			Data:    activity,
		})

		res.SetBody(data)
		res.SetStatusCode(201)
		return nil
	})

	// get all activity group
	router.Get("/activity-groups", func(ctx *fiber.Ctx) error {
		res := ctx.Response()
		res.Header.Set("Content-Type", "application/json")

		listActivity, err := usecase.GetAll()
		if err != nil {
			return err
		}

		data, _ := json.Marshal(dtos.Response{
			Status:  "success",
			Message: "success to get list activity group",
			Data:    listActivity,
		})

		res.SetBody(data)
		res.SetStatusCode(200)
		return nil
	})

	// get activity by id
	router.Get("/activity-groups/:id", func(ctx *fiber.Ctx) error {
		res := ctx.Response()
		res.Header.Set("Content-Type", "application/json")

		id := ctx.Params("id", "0")
		activity, err := usecase.GetById(id)
		if err != nil {
			return err
		}

		if activity.Title == "" {
			data, _ := json.Marshal(dtos.Response{
				Status:  "failed",
				Message: "activity not found",
			})

			res.SetBody(data)
			res.SetStatusCode(404)

			return nil
		}

		data, _ := json.Marshal(dtos.Response{
			Status:  "success",
			Message: "success to get activity",
			Data:    activity,
		})

		res.SetBody(data)
		res.SetStatusCode(200)
		return nil
	})

	// delete activity by id
	router.Delete("/activity-groups/:id", func(ctx *fiber.Ctx) error {
		res := ctx.Response()
		res.Header.Set("Content-Type", "application/json")

		id := ctx.Params("id", "-1")
		isDeleted, err := usecase.DeleteById(id)
		if err != nil {
			return err
		}

		if !isDeleted {
			data, _ := json.Marshal(dtos.Response{
				Status:  "failed",
				Message: "activity not found",
			})

			res.SetBody(data)
			res.SetStatusCode(404)

			return nil
		}

		data, _ := json.Marshal(dtos.Response{
			Status:  "success",
			Message: "success to delete activity",
		})

		res.SetBody(data)
		res.SetStatusCode(200)
		return nil
	})

	// update activity group
	router.Patch("/activity-groups/:id", func(ctx *fiber.Ctx) error {
		res := ctx.Response()
		res.Header.Set("Content-Type", "application/json")

		var req *dtos.ActivityUpdateRequest
		err := ctx.BodyParser(&req)
		if err != nil {
			// failed to parse request body
			data, _ := json.Marshal(dtos.Response{
				Status:  "error",
				Message: "can not parse request body",
			})

			res.SetBody(data)
			res.SetStatusCode(400)

			return nil
		}

		err = validate.Struct(req)
		if err != nil {
			// failed to validate request body
			data, _ := json.Marshal(dtos.Response{
				Status:  "failed",
				Message: err.Error(),
			})

			res.SetBody(data)
			res.SetStatusCode(401)

			return nil
		}

		id := ctx.Params("id", "-1")
		isUpdated, activity, _ := usecase.UpdateById(id, req)
		if !isUpdated {
			data, _ := json.Marshal(dtos.Response{
				Status:  "failed",
				Message: "activity not found",
			})

			res.SetBody(data)
			res.SetStatusCode(404)
			return nil
		}

		data, _ := json.Marshal(dtos.Response{
			Status:  "success",
			Message: "success to update activity group",
			Data:    activity,
		})

		res.SetBody(data)
		res.SetStatusCode(200)
		return nil
	})
}
