package users

import "time"

type (
	UserModel struct {
		ID        int64     `db:"id"`
		Email     string    `db:"email"`
		Password  string    `db:"password"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

type (
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

type (
	CreateUserResponse struct {
		ID    int64  `json:"id"`
		Email string `json:"email"`
	}
)
