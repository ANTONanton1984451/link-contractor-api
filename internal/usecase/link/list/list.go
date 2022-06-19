package list

import (
	"context"
	"fmt"
	"link-contractor-api/internal/entities/user"
	"time"
)

type (
	List interface {
		Execute(ctx context.Context, usr user.User, option SelectOption) ([]Link, error)
	}

	SelectOption struct {
		OnlyActive bool
	}

	Link struct {
		Path       string
		RedirectTo string
		Active     bool
		CreatedAt  time.Time
	}

	LinkRepo interface {
		ListLinks(ctx context.Context, usr user.User, option SelectOption) ([]Link, error)
	}
)

type usecase struct {
	linkRepo LinkRepo
}

func New(linkRepo LinkRepo) List {
	return &usecase{linkRepo: linkRepo}
}

func (uc *usecase) Execute(ctx context.Context, usr user.User, option SelectOption) ([]Link, error) {
	linkList, err := uc.linkRepo.ListLinks(ctx, usr, option)
	if err != nil {
		return nil, fmt.Errorf("list links: %w", err)
	}

	return linkList, nil
}
