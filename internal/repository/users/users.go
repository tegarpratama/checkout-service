package users

import (
	"context"
	"database/sql"

	"github.com/tegarpratama/checkout-service/internal/model/users"
)

func (r *repository) CreateUser(ctx context.Context, model users.UserModel) (int64, error) {
	query := `INSERT INTO users (email, password, created_at, updated_at) VALUES (?, ?, ?, ?)`
	result, err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.CreatedAt, model.UpdatedAt)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *repository) GetUser(ctx context.Context, email string) (*users.UserModel, error) {
	var response users.UserModel

	query := `SELECT id FROM users WHERE email = ?`
	row := r.db.QueryRowContext(ctx, query, email)
	err := row.Scan(&response.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &response, nil
}
