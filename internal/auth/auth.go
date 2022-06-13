package auth

import (
	"context"
	"fmt"
	"link-contractor-api/internal/entrypoint"
	"time"
)

type (
	BanList interface {
		InBan(ctx context.Context, user entrypoint.User) (bool, *time.Time, error)
	}

	UserRepo interface {
		// если юзера не было - зарегать его
		GetByExternalID(ctx context.Context, externalID int64) (entrypoint.User, error)
	}
)

type auth struct {
	banList  BanList
	userRepo UserRepo
}

func (a *auth) GetUser(ctx context.Context, externalID int64) (entrypoint.User, error) {
	user, err := a.userRepo.GetByExternalID(ctx, externalID)
	if err != nil {
		return entrypoint.User{}, fmt.Errorf("get user by externalID: %w", err)
	}

	inBan, until, err := a.banList.InBan(ctx, user)
	if err != nil {
		return entrypoint.User{}, fmt.Errorf("check user in ban list: %w", err)
	}

	if inBan {
		return entrypoint.User{}, entrypoint.UserInBanErr{Until: until}
	}

	return user, err
}

func New(bl BanList, uRepo UserRepo) *auth {
	return &auth{
		banList:  bl,
		userRepo: uRepo,
	}
}
