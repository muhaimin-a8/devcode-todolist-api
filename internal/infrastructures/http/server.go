package http

import (
	"context"
	"devcode-todolist-api/internal/applications/usecase"
	"devcode-todolist-api/internal/infrastructures/database"
	"devcode-todolist-api/internal/infrastructures/repository"
	"devcode-todolist-api/internal/interfaces/http/api"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"net"
	"os"
)

type Server struct {
	app *fiber.App
}

//user:password@tcp(localhost:3306)/your_database
func CreateServer(ctx context.Context) (*Server, error) {
	// establish database connection
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	db, err := database.NewDBMySQL(dsn)
	if err != nil {
		log.Fatalln(err)
	}

	// validator
	validate := validator.New()

	server := &Server{}
	server.app = fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Println(err)
			return err
		},
	})

	// repositories
	activityRepository := repository.NewActivityRepositoryMySQL(db)
	todoRepository := repository.NewTodoRepositoryMySQL(db)

	// error handling
	//server.app.Use(recover.New())

	// usecase
	activityUseCase := usecase.NewActivityUseCase(activityRepository)
	todoUseCase := usecase.NewTodoUseCase(todoRepository, activityRepository)

	//register router
	api.RegisterActivityController(server.app.Group("/activity-groups"), validate, activityUseCase)
	api.RegisterTodoController(server.app.Group("/todo-items"), validate, todoUseCase)

	return server, nil
}

func (s *Server) Serve(listener net.Listener) error {
	return s.app.Listener(listener)
}
