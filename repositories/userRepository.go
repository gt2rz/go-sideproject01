package repositories

import (
	"context"
	"database/sql"
	"errors"

	"microtwo/models"
)

var ErrUserNotFound = errors.New("user not found")
var ErrUserNotSaved = errors.New("user not saved")

type UserRepository interface {
	SaveUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, id string) (*models.User, error)
	GenerateResetToken(ctx context.Context, id string) (string, error)
	GetUserByResetToken(ctx context.Context, resetToken string) (*models.User, error)
	UpdatePassword(ctx context.Context, id string, password string) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) (*UserRepositoryImpl, error) {
	return &UserRepositoryImpl{db}, nil
}

func (r *UserRepositoryImpl) SaveUser(ctx context.Context, user models.User) error {

	result, err := r.db.ExecContext(ctx, "INSERT INTO users (id, email, password, first_name, last_name, phone, verified, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?,?, ?, ?)", user.Id, user.Email, user.Password, user.Firstname, user.Lastname, user.Phone, user.Verified, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return ErrUserNotSaved
	}
	return nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	var user = models.User{}

	query := r.db.QueryRowContext(ctx, "SELECT id, email, password, first_name, last_name, phone, verified, created_at, updated_at  FROM users WHERE email=?", email)
	err := query.Scan(&user.Id, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Phone, &user.Verified, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *UserRepositoryImpl) GenerateResetToken(ctx context.Context, id string) (string, error) {

	return "", nil
}

func (r *UserRepositoryImpl) GetUserByResetToken(ctx context.Context, resetToken string) (*models.User, error) {
	var user = models.User{}

	query := r.db.QueryRowContext(ctx, "SELECT id, email, first_name  FROM users WHERE password_reset_token=?", resetToken)
	err := query.Scan(&user.Id, &user.Email, &user.Firstname)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) UpdatePassword(ctx context.Context, id string, password string) error {
	result, err := r.db.ExecContext(ctx, "UPDATE users SET password=? WHERE id=?", password, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return ErrUserNotSaved
	}
	return nil
}
