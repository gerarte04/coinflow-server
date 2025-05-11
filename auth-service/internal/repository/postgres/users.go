package postgres

import (
	"coinflow/coinflow-server/auth-service/internal/models"
	"coinflow/coinflow-server/auth-service/internal/repository"
	"coinflow/coinflow-server/pkg/database"
	"coinflow/coinflow-server/pkg/database/postgres"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepo struct {
	conn *pgx.Conn
}

func NewUsersRepo(conn *pgx.Conn) *UsersRepo {
	return &UsersRepo{conn: conn}
}

func (r *UsersRepo) GetUser(id uuid.UUID) (*models.User, error) {
	const op = "UsersRepo.GetUser"

	row := r.conn.QueryRow(
		context.Background(),
		"SELECT id, login, name, email, phone, register_timestamp FROM users WHERE id = $1",
		id,
	)

	var usr models.User
	err := row.Scan(&usr.Id, &usr.Login, &usr.Name, &usr.Email, &usr.Phone, &usr.RegistrationTimestamp)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrorUserIdNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &usr, nil
}

func (r *UsersRepo) GetUserByCred(login, password string) (*models.User, error) {
	const op = "UsersRepo.GetUserByCred"

	row := r.conn.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE login = $1",
		login,
	)

	var usr models.User
	passwordHash := make([]byte, 0)

	err := row.Scan(&usr.Id, &usr.Login, &passwordHash, &usr.Name, &usr.Email, &usr.Phone, &usr.RegistrationTimestamp)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrorUserLoginNotFound)
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, fmt.Errorf("%s: %w", op, repository.ErrorWrongPassword)
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &usr, nil
}

func (r *UsersRepo) PostUser(usr *models.User) (uuid.UUID, error) {
	const op = "UsersRepo.PostUser"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 16)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	row := r.conn.QueryRow(
		context.Background(),
		`INSERT INTO users (
			login, password_hash, name, email, phone
		) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		usr.Login, passwordHash, usr.Name, usr.Email, usr.Phone,
	)

	var usrId uuid.UUID
	err = row.Scan(&usrId)

	if dbErr := postgres.DetectError(err); dbErr == database.ErrorUniqueViolation {
		return uuid.Nil, fmt.Errorf("%s: %w", op, repository.ErrorUserCredAlreadyExists)
	} else if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return usrId, nil
}
