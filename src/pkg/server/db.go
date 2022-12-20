package server

import (
	"database/sql"

	"github.com/carlosghabrous/never-red/src/pkg/models"
)

type dbDriver struct {
	driver *sql.DB
}

type DbQuerier interface {
	getAll() []models.Movement
	getMovement(id int) models.Movement
}

func (db dbDriver) getAll() ([]models.Movement, error) {
	result := []models.Movement{}
	return result, nil
}

func (db dbDriver) getMovement(id int) (models.Movement, error) {
	result := models.Movement{}
	return result, nil
}
