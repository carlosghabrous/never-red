package server

import (
	"database/sql"

	"github.com/carlosghabrous/never-red/src/pkg/models"
)

type dbDriver struct {
	driver *sql.DB
}

type DbQuerier interface {
	getMovements() ([]models.Movement, error)
	getMovement(id int) (models.Movement, error)
}

func (db dbDriver) getMovements() ([]models.Movement, error) {
	result := []models.Movement{}
	return result, nil
}

func (db dbDriver) getMovement(id int) (models.Movement, error) {
	result := models.Movement{}
	return result, nil
}
