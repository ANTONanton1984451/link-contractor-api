package activatepath

import (
	"context"
	"errors"
	"fmt"
)

type (
	ActivatePath interface {
		Execute(ctx context.Context, path Path, user User) error
	}

	Path struct {
		Path string
	}

	User struct {
		ID int64
	}

	LinkRepo interface {
		UserOwnThisPath(ctx context.Context, path string, userID int64) (bool, error)
		ActivatePath(ctx context.Context, path string) error
	}
)

var (
	UserIsNotOwnerOfPathErr = errors.New("user dont own this path")
	PathDontExistErr        = errors.New("path dont exist")
)

type usecase struct {
	linkRepo LinkRepo
}

func New(linkRepo LinkRepo) ActivatePath {
	return &usecase{linkRepo: linkRepo}
}

func (uc *usecase) Execute(ctx context.Context, path Path, user User) error {
	isOwn, err := uc.linkRepo.UserOwnThisPath(ctx, path.Path, user.ID)
	if err != nil {
		return fmt.Errorf("chek is owner user: %w", err)
	}
	if !isOwn {
		return UserIsNotOwnerOfPathErr
	}

	err = uc.linkRepo.ActivatePath(ctx, path.Path)
	if err != nil {
		return fmt.Errorf("activate path: %w", err)
	}

	return nil
}
