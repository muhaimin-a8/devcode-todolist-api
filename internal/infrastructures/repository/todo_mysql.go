package repository

import (
	"database/sql"
	"devcode-todolist-api/internal/domains"
)

type todoRepositoryMySQLImpl struct {
	DB *sql.DB
}

func (t todoRepositoryMySQLImpl) Update(todo domains.Todo) (*domains.Todo, error) {
	stmt, err := t.DB.Prepare("UPDATE todos SET title = ?, priority = ? WHERE todo_id = ?")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(todo.Title, todo.Priority, todo.Id)

	rows, _ := t.DB.Query("SELECT * FROM todos WHERE todo_id = ?", todo.Id)

	var todoFromDB domains.Todo
	for rows.Next() {
		err = rows.Scan(&todoFromDB.Id, &todoFromDB.ActivityGroupId, &todoFromDB.Title, &todoFromDB.Priority, &todoFromDB.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &todoFromDB, nil
}

func (t todoRepositoryMySQLImpl) DeleteById(id string) (isDeleted bool, err error) {
	res, err := t.DB.Exec("DELETE FROM todos WHERE todo_id = ?", id)
	if err != nil {
		return false, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected != 1 {
		return false, nil
	}

	return true, nil
}

func (t todoRepositoryMySQLImpl) GetById(id string) (*domains.Todo, error) {
	rows, err := t.DB.Query("SELECT * FROM todos WHERE todo_id = ?", id)
	if err != nil {
		return nil, err
	}

	var todoFromDB domains.Todo
	for rows.Next() {
		err := rows.Scan(&todoFromDB.Id, &todoFromDB.ActivityGroupId, &todoFromDB.Title, &todoFromDB.Priority, &todoFromDB.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &todoFromDB, err
}

func (t todoRepositoryMySQLImpl) GetAllByActivityId(activityId string) ([]domains.Todo, error) {
	rows, err := t.DB.Query("SELECT * FROM todos WHERE activity_group_id = ?", activityId)
	if err != nil {
		return nil, err
	}

	var todoFromDB domains.Todo
	var listTodo []domains.Todo
	for rows.Next() {
		err := rows.Scan(&todoFromDB.Id, &todoFromDB.ActivityGroupId, &todoFromDB.Title, &todoFromDB.Priority, &todoFromDB.CreatedAt)
		if err != nil {
			return nil, err
		}

		listTodo = append(listTodo, todoFromDB)
	}

	return listTodo, err
}

func (t todoRepositoryMySQLImpl) Save(todo domains.Todo) (*domains.Todo, error) {
	stmt, err := t.DB.Prepare("INSERT INTO todos (activity_group_id, title, priority) VALUES(?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(todo.ActivityGroupId, todo.Title, todo.Priority)

	rows, _ := t.DB.Query("SELECT * FROM todos WHERE todo_id = LAST_INSERT_ID()")

	var todoFromDB domains.Todo
	for rows.Next() {
		err = rows.Scan(&todoFromDB.Id, &todoFromDB.ActivityGroupId, &todoFromDB.Title, &todoFromDB.Priority, &todoFromDB.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &todoFromDB, nil
}

func NewTodoRepositoryMySQL(db *sql.DB) domains.TodoRepository {
	return &todoRepositoryMySQLImpl{DB: db}
}
