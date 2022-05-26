package list

import (
	"context"
	"fmt"
	"time"
)

type (
	List interface {
		Execute(ctx context.Context, option SelectOption) ([]Link, error)
	}

	SelectOption struct {
		WithActivate bool
	}

	Link struct {
		Path       string
		RedirectTo string
		Active     bool
		CreatedAt  time.Time
	}

	LinkRepo interface {
		ListLinks(ctx context.Context, option SelectOption) ([]Link, error)
	}
)

type usecase struct {
	linkRepo LinkRepo
}

func (uc *usecase) Execute(ctx context.Context, option SelectOption) ([]Link, error) {
	linkList, err := uc.linkRepo.ListLinks(ctx, option)
	if err != nil {
		return nil, fmt.Errorf("list links: %w", err)
	}

	return linkList, nil
}
