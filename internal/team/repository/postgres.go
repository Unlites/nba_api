package repository

import (
	"fmt"

	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/team"
	"github.com/jmoiron/sqlx"
)

type teamRepo struct {
	db *sqlx.DB
}

func NewTeamRepo(db *sqlx.DB) team.Repository {
	return &teamRepo{db: db}
}

func (r *teamRepo) GetById(id int64) (*models.Team, error) {
	team := &models.Team{}
	query := fmt.Sprintf(selectTeamByIdQuery, teamsTable)
	err := r.db.Get(team, query, id)

	return team, err
}

func (r *teamRepo) Create(team *models.Team) error {
	query := fmt.Sprintf(insertTeamQuery, teamsTable)
	_, err := r.db.Exec(query, team.Name)

	return err
}

func (r *teamRepo) Update(team *models.Team) error {
	query := fmt.Sprintf(updateTeamQuery, teamsTable)
	_, err := r.db.Exec(query, team.Name)

	return err
}

func (r *teamRepo) Delete(id int64) error {
	query := fmt.Sprintf(deleteTeamQuery, teamsTable)
	_, err := r.db.Exec(query, id)

	return err
}
