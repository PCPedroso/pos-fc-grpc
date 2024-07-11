package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name string, description string, categoryID string) (*Course, error) {
	id := uuid.New().String()

	stmt, err := c.db.Prepare("INSERT INTO courses (id, name, description, category_id) VALUES (?,?,?,?);")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, description, categoryID)
	if err != nil {
		return nil, err
	}
	stmt.Close()

	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	coursess := []Course{}
	for rows.Next() {
		var id, name, description, category_id string
		if err := rows.Scan(&id, &name, &description, &category_id); err != nil {
			return nil, err
		}
		coursess = append(coursess, Course{ID: id, Name: name, Description: description, CategoryID: category_id})
	}
	return coursess, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	stmt, err := c.db.Prepare("SELECT id, name, description, category_id FROM courses where category_id = ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(categoryID)
	if err != nil {
		return nil, err
	}
	stmt.Close()

	coursess := []Course{}
	for rows.Next() {
		var id, name, description, category_id string
		if err := rows.Scan(&id, &name, &description, &category_id); err != nil {
			return nil, err
		}
		coursess = append(coursess, Course{ID: id, Name: name, Description: description, CategoryID: category_id})
	}
	return coursess, nil
}
