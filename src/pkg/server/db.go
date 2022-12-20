package server

import (
	"database/sql"

	"github.com/carlosghabrous/never-red/src/pkg/models"
)

type dbDriver struct {
	db *sql.DB
}

type DbQuerier interface {
	getAll() []models.Movement
	getMovement(id int) models.Movement
}

func (dbDriver) getAll() ([]models.Movement, error) {
	result := []models.Movement{}
	return result, nil
}

func (dbDriver) getMovement(id int) (models.Movement, error) {
	result := models.Movement{}
	return result, nil
}
