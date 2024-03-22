package api

import (
	"devcode-todolist-api/internal/domains"
	"devcode-todolist-api/internal/dtos"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func RegisterTodoController(router fiber.Router, validate *validator.Validate, usecase domains.TodoUseCase) {
	// create new todos
	router.Post("/", func(ctx *fiber.Ctx) error {
		res := ctx.Response()
		res.Header.Set("Content-Type", "application/json")

		var req *dtos.TodoCreateRequest
		err := ctx.BodyParser(&req)
		if err != nil {
			// failed to parse request body
			data, _ := json.Marshal(dtos.Response{
				Status:  "error",
				Message: "can not parse request body",
			})

			res.SetBody(data)
			res.SetStatusCode(404)

			return nil
		}

		err = validate.Struct(req)
		if err != nil {
			// failed to validate request body
			response := dtos.Response{
				Status:  "Bad Request",
				Message: "activity_group_id cannot be null",
			}

			if req.Title == "" {
				response.Message = "title cannot be null"
			}

			data, _ := json.Marshal(response)

			res.SetBody(data)
			res.SetStatusCode(400)

			return nil
		}

		isActivityExist, todo, err := usecase.CreateNew(req)
		if !isActivityExist {
			data, _ := json.Marshal(dtos.Response{
				Status:  "failed",
				Message: "activity not found",
			})

			res.SetBody(data)
			res.SetStatusCode(404)
			return nil
		}

		data, _ := json.Marshal(dtos.Response{
			Status:  "Success",
			Message: "success to create new todo",
			Data:    todo,
		})

		res.SetBody(data)
		res.SetStatusCode(201)
		return nil
	})

	// get all todos by activity group id
	router.Get("/", func(ctx *fiber.Ctx) error {
		res := ctx.Response()
		res.Header.Set("Content-Type", "application/json")

		activityId, _ := strconv.Atoi(ctx.Query("activity_group_id", "-1"))

		_, todos, err := usecase.GetAllByActivityId(activityId)
		if err != nil {
			return err
		}
		//if !isActivityExist {
		//	data, _ := json.Marshal(dtos.Response{
		//		Status:  "failed",
		//		Message: "activity not found",
		//	})
		//
		//	res.SetBody(data)
		//	res.SetStatusCode(404)
		//	return nil
		//}

		data, _ := json.Marshal(dtos.Response{
			Status:  "Success",
			Message: "success to get list todo",
			Data:    todos,
		})

		res.SetBody(data)
		res.SetStatusCode(200)
		return nil
	})

	// get todos by id
	router.Get("/:id", func(ctx *fiber.Ctx) error {
		res := ctx.Response()
		res.Header.Set("Content-Type", "application/json")

		id, _ := strconv.Atoi(ctx.Params("id", "-1"))
		todo, err := usecase.GetById(id)
		if err != nil {
			return err
		}

		if todo.Title == "" {
			data, _ := json.Marshal(dtos.Response{
				Status:  "Not Found",
				Message: fmt.Sprintf("Todo with ID %d Not Found", id),
			})

			res.SetBody(data)
			res.SetStatusCode(404)

			return nil
		}

		data, _ := json.Marshal(dtos.Response{
			Status:  "Success",
			Message: "success to get todos",
			Data:    todo,
		})

		res.SetBody(data)
		res.SetStatusCode(200)
		return nil
	})

	// delete todos by id
	router.Delete("/:id", func(ctx *fiber.Ctx) error {
		res := ctx.Response()
		res.Header.Set("Content-Type", "application/json")

		id, _ := strconv.Atoi(ctx.Params("id", "-1"))
		isDeleted, err := usecase.DeleteById(id)
		if err != nil {
			return err
		}

		if !isDeleted {
			data, _ := json.Marshal(dtos.Response{
				Status:  "Not Found",
				Message: fmt.Sprintf("Todo with ID %d Not Found", id),
			})

			res.SetBody(data)
			res.SetStatusCode(404)

			return nil
		}

		data, _ := json.Marshal(dtos.Response{
			Status:  "Success",
			Message: "success to delete todos",
			Data:    struct{}{},
		})

		res.SetBody(data)
		res.SetStatusCode(200)
		return nil
	})

	// update todos
	router.Patch("/:id", func(ctx *fiber.Ctx) error {
		res := ctx.Response()
		res.Header.Set("Content-Type", "application/json")

		var req *dtos.TodoUpdateRequest
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
			res.SetStatusCode(400)

			return nil
		}

		id, _ := strconv.Atoi(ctx.Params("id", "-1"))
		isUpdated, todo, _ := usecase.UpdateById(id, req)
		if !isUpdated {
			data, _ := json.Marshal(dtos.Response{
				Status:  "Not Found",
				Message: fmt.Sprintf("Todo with ID %d Not Found", id),
			})

			res.SetBody(data)
			res.SetStatusCode(404)
			return nil
		}

		data, _ := json.Marshal(dtos.Response{
			Status:  "Success",
			Message: "success to update todo",
			Data:    todo,
		})

		res.SetBody(data)
		res.SetStatusCode(200)
		return nil
	})
}
