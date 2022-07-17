package auth

import (
	"context"
	"fmt"
	"link-contractor-api/internal/entities/user"
	"link-contractor-api/internal/entrypoint"
	"time"
)

type (
	BanList interface {
		InBan(ctx context.Context, user user.User) (bool, *time.Time, error)
	}

	UserRepo interface {
		GetByExternalID(ctx context.Context, user user.User) (user.User, error)
	}
)

type auth struct {
	banList  BanList
	userRepo UserRepo
}

// GetUser получаем пользователя по external id и поверяем, в бане ли он сейчас
func (a *auth) GetUser(ctx context.Context, usr user.User) (user.User, error) {
	usr, err := a.userRepo.GetByExternalID(ctx, usr)
	if err != nil {
		return user.User{}, fmt.Errorf("get user by externalID: %w", err)
	}

	inBan, until, err := a.banList.InBan(ctx, usr)
	if err != nil {
		return user.User{}, fmt.Errorf("check user in ban list: %w", err)
	}

	if inBan {
		return user.User{}, entrypoint.UserInBanErr{Until: until}
	}

	return usr, err
}

func New(bl BanList, uRepo UserRepo) *auth {
	return &auth{
		banList:  bl,
		userRepo: uRepo,
	}
}
