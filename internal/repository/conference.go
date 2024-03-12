package repository

import (
	"clean-API/internal/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type ConferenceRepository interface {
	Save(conference model.Conference) (model.Conference, error)
	GetAllConferences() ([]model.Conference, error)
	GetConferenceByID(id int64) (model.Conference, error)
	Update(conference model.Conference) error
	Delete(conference model.Conference) error
}

type conference struct {
	db *sql.DB
}

func newConferenceRepository(db *sql.DB) ConferenceRepository {
	return &conference{
		db: db,
	}
}

func (c conference) Save(conference model.Conference) (model.Conference, error) {
	query := `INSERT INTO conferences(name, description, location, dateTime, user_id) VALUES(?, ?, ?, ?, ?)`

	stmt, err := c.db.Prepare(query)

	if err != nil {
		return model.Conference{}, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}(stmt)

	result, err := stmt.Exec(conference.Name, conference.Description, conference.Location,
		conference.DateTime, conference.UserID)

	if err != nil {
		return model.Conference{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return model.Conference{}, err
	}
	conference.ID = id
	return conference, nil
}

func (c conference) GetAllConferences() ([]model.Conference, error) {
	query := "SELECT * FROM conferences"

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	var conferences []model.Conference

	for rows.Next() {
		var conference model.Conference
		err := rows.Scan(&conference.ID, &conference.Name, &conference.Description, &conference.Location, &conference.DateTime, &conference.UserID)

		if err != nil {
			return nil, err
		}
		conferences = append(conferences, conference)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return conferences, nil

}

func (c conference) GetConferenceByID(id int64) (model.Conference, error) {
	query := "SELECT * FROM conferences WHERE id = ?"
	row := c.db.QueryRow(query, id)

	var conference model.Conference
	err := row.Scan(&conference.ID, &conference.Name, &conference.Description, &conference.Location, &conference.DateTime, &conference.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Conference{}, fmt.Errorf("conference with ID %d not found", id)
		} else {
			log.Fatal(err)
		}
	}
	return conference, nil
}

func (c conference) Update(conference model.Conference) error {
	query := `
	UPDATE conferences
	SET name = ?, description = ?, location = ?, dateTime = ? 
	WHERE id = ?
	`

	stmt, err := c.db.Prepare(query)

	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}(stmt)

	_, err = stmt.Exec(conference.Name, conference.Description, conference.Location, conference.DateTime, conference.ID)
	return err
}

func (c conference) Delete(conference model.Conference) error {
	query := `DELETE FROM conferences WHERE id = ?`

	stmt, err := c.db.Prepare(query)

	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}(stmt)

	_, err = stmt.Exec(conference.ID)
	return err
}
