package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	userEntity "link-contractor-api/internal/entities/user"
	"time"

	"github.com/jackc/pgx"
)

// GetByExternalID имплементация интерфейса для компонента auth
func (repo *repository) GetByExternalID(ctx context.Context, user userEntity.User) (userEntity.User, error) {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return userEntity.User{}, fmt.Errorf("get connection: %w", err)
	}

	defer conn.Close()

	querySource := entrypointInnerSource[user.SourceType]

	var userR userRow
	q, a := getUserQuery(user.ExternalID, querySource)

	err = conn.QueryRow(q, a...).Scan(&userR.ID, &userR.Name, &userR.Surname, &userR.SourceID, &userR.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repo.createUser(ctx, user)
		}
		return userEntity.User{}, fmt.Errorf("exect query: %w", err)
	}

	userR.ExternalID = user.ExternalID

	userEnt := mapUserToEntityUser(userR)
	userEnt.SourceType = user.SourceType

	return userEnt, nil
}

func (repo *repository) createUser(ctx context.Context, user userEntity.User) (userEntity.User, error) {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return userEntity.User{}, fmt.Errorf("get connection: %w", err)
	}

	defer conn.Close()

	var insertID int64
	q, a := createUserQuery(user)
	err = conn.QueryRow(q, a...).Scan(&insertID)
	if err != nil {
		return userEntity.User{}, fmt.Errorf("exec query: %w", err)
	}

	user.ID = insertID

	return user, nil
}

// InBan имплементация интерфейса банлиста для компонента auth
func (repo *repository) InBan(ctx context.Context, user userEntity.User) (bool, *time.Time, error) {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return false, nil, fmt.Errorf("get connection: %w", err)
	}
	defer conn.Close()

	var until sql.NullTime

	q, a := getUserFromBanQuery(user)
	err = conn.QueryRow(q, a...).Scan(&until)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil, nil
		}
		return false, nil, fmt.Errorf("exec query: %w", err)
	}

	if until.Valid && until.Time.Before(time.Now()) {
		return false, nil, nil
	}

	if !until.Valid {
		return true, nil, nil
	}
	return true, &until.Time, nil
}

func mapUserToEntityUser(uRow userRow) userEntity.User {
	return userEntity.User{
		ID:           uRow.ID,
		Name:         uRow.Name,
		Surname:      uRow.Surname,
		ExternalID:   uRow.ExternalID,
		RegisteredAt: uRow.CreatedAt,
	}
}
