package repository

import (
	"database/sql"
	"devcode-todolist-api/internal/domains"
)

type activityRepositoryMySQLImpl struct {
	DB *sql.DB
}

func (a activityRepositoryMySQLImpl) Update(activity domains.Activity) (*domains.Activity, error) {
	stmt, err := a.DB.Prepare("UPDATE activities SET title = ?, email = ? WHERE activity_id = ?")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(activity.Title, activity.Email, activity.Id)

	rows, _ := a.DB.Query("SELECT * FROM activities WHERE activity_id = ?", activity.Id)

	var activityFromDB domains.Activity
	for rows.Next() {
		err = rows.Scan(&activityFromDB.Id, &activityFromDB.Title, &activityFromDB.Email, &activityFromDB.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &activityFromDB, nil
}

func (a activityRepositoryMySQLImpl) DeleteById(id int) (bool, error) {
	res, err := a.DB.Exec("DELETE FROM activities WHERE activity_id = ?", id)
	if err != nil {
		return false, err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected != 1 {
		return false, nil
	}

	return true, nil
}

func (a activityRepositoryMySQLImpl) GetById(id int) (*domains.Activity, error) {
	stmt, err := a.DB.Prepare("SELECT * FROM activities WHERE activity_id = ?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)

	var activityFromDB domains.Activity
	for rows.Next() {
		err = rows.Scan(&activityFromDB.Id, &activityFromDB.Title, &activityFromDB.Email, &activityFromDB.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &activityFromDB, nil
}

func (a activityRepositoryMySQLImpl) GetAll() ([]domains.Activity, error) {
	rows, err := a.DB.Query("SELECT * FROM activities")
	if err != nil {
		return nil, err
	}

	var activityFromDB domains.Activity
	var listActivity []domains.Activity
	for rows.Next() {
		err := rows.Scan(&activityFromDB.Id, &activityFromDB.Title, &activityFromDB.Email, &activityFromDB.CreatedAt)
		if err != nil {
			return nil, err
		}

		listActivity = append(listActivity, activityFromDB)
	}

	return listActivity, err
}

func (a activityRepositoryMySQLImpl) Save(activity domains.Activity) (*domains.Activity, error) {
	stmt, err := a.DB.Prepare("INSERT INTO activities (title, email) VALUES(?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(activity.Title, activity.Email)

	rows, _ := a.DB.Query("SELECT * FROM activities WHERE activity_id = LAST_INSERT_ID()")

	var activityFromDB domains.Activity
	for rows.Next() {
		err = rows.Scan(&activityFromDB.Id, &activityFromDB.Title, &activityFromDB.Email, &activityFromDB.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &activityFromDB, nil
}

func NewActivityRepositoryMySQL(db *sql.DB) domains.ActivityRepository {
	return &activityRepositoryMySQLImpl{DB: db}
}
