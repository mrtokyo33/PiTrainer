package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/mrtokyo33/PiTrainer/src/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO users (username, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
}
