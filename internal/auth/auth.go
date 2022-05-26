package auth

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type (
	Auth interface {
		GetUser(ctx context.Context, externalID int64) (User, error)
	}

	BanList interface {
		InBan(ctx context.Context, user User) (bool, error)
	}

	UserRepo interface {
		GetByExternalID(ctx context.Context, externalID int64) (User, error)
	}

	User struct {
		Name         string
		SurName      string
		ID           int64
		ExternalID   int64
		RegisteredAt time.Time
	}
)

var (
	UserInBanErr = errors.New("user in ban")
)

type auth struct {
	banList  BanList
	userRepo UserRepo
}

func (a *auth) GetUser(ctx context.Context, externalID int64) (User, error) {
	user, err := a.userRepo.GetByExternalID(ctx, externalID)
	if err != nil {
		return User{}, fmt.Errorf("get user by externalID: %w", err)
	}

	inBan, err := a.banList.InBan(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("check user in ban list: %w", err)
	}

	if inBan {
		return User{}, UserInBanErr
	}

	return user, err
}