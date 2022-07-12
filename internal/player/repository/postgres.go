package repository

import (
	"fmt"

	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/player"
	"github.com/jmoiron/sqlx"
)

type playerRepo struct {
	db *sqlx.DB
}

func NewPlayerRepo(db *sqlx.DB) player.Repository {
	return &playerRepo{db: db}
}

func (r *playerRepo) GetById(id int64) (*models.Player, error) {
	var player *models.Player
	query := fmt.Sprintf(selectPlayerByIdQuery, playersTable)
	err := r.db.Get(&player, query, id)

	return player, err
}

func (r *playerRepo) GetAllByTeamId(teamId int64) ([]*models.Player, error) {
	var players []*models.Player
	query := fmt.Sprintf(selectStatByTeamIdQuery, playersTable)
	err := r.db.Select(&players, query, teamId)

	return players, err
}

func (r *playerRepo) Create(player *models.Player) error {
	query := fmt.Sprintf(insertPlayerQuery, playersTable)
	_, err := r.db.Exec(query, player.Name, player.TeamId)

	return err
}
