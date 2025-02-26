package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (*Category, error) {
	id := uuid.New().String()

	stmt, err := c.db.Prepare("INSERT INTO categories (id, name, description) VALUES (?,?,?);")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, description)
	if err != nil {
		return nil, err
	}

	return &Category{
		ID:          id,
		Name:        name,
		Description: description,
	}, nil

}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}
	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var id, name, description string
	stmt, err := c.db.Prepare("SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id WHERE co.id = ?;")
	if err != nil {
		return Category{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(courseID).Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindByCategoryID(categoryID string) (Category, error) {
	var id, name, description string
	stmt, err := c.db.Prepare("SELECT id, name, description FROM categories WHERE id = ?;")
	if err != nil {
		return Category{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(categoryID).Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}
