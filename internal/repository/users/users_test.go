package users

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/tegarpratama/checkout-service/internal/model/users"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	repo := NewRepository(db)

	model := users.UserModel{
		Email:     "test@example.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectExec(`INSERT INTO users`).
		WithArgs(model.Email, model.Password, model.CreatedAt, model.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	userID, err := repo.CreateUser(context.Background(), model)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), userID)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	repo := NewRepository(db)

	expectedUser := &users.UserModel{
		ID:    1,
		Email: "test@example.com",
	}

	mock.ExpectQuery(`SELECT id FROM users WHERE email = ?`).
		WithArgs(expectedUser.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedUser.ID))

	user, err := repo.GetUser(context.Background(), expectedUser.Email)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.ID, user.ID)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetUser_NoRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	repo := NewRepository(db)

	mock.ExpectQuery(`SELECT id FROM users WHERE email = ?`).
		WithArgs("nonuser@example.com").
		WillReturnError(sql.ErrNoRows)

	user, err := repo.GetUser(context.Background(), "nonuser@example.com")

	assert.NoError(t, err)
	assert.Nil(t, user)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
