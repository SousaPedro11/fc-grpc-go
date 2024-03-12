package database

import (
	"database/sql"
	"errors"
	"fmt"

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

	if name == "" || categoryID == "" {
		return nil, errors.New("name and categoryID are required")
	}

	id := uuid.New().String()

	stmt, err := c.db.Prepare("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, name, description, categoryID)
	if err != nil {
		return nil, err
	}

	return &Course{ID: id, Name: name, Description: description, CategoryID: categoryID}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, fmt.Errorf("error querying courses: %w", err)
	}
	defer rows.Close()

	courses := make([]Course, 0)
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, fmt.Errorf("error scanning course: %w", err)
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading courses: %w", err)
	}

	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, fmt.Errorf("error querying courses: %w", err)
	}
	defer rows.Close()

	courses := make([]Course, 0)
	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if err != nil {
			return nil, fmt.Errorf("error scanning course: %w", err)
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading courses: %w", err)
	}

	return courses, nil
}
