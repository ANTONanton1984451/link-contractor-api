package link

import (
	"context"
	"errors"
	"fmt"
	"link-contractor-api/internal/usecase/link/create"
	"link-contractor-api/internal/usecase/link/list"

	"github.com/jackc/pgx"
)

func (repo *repository) UserHasThisLink(ctx context.Context, linkToCheck string, userID int64) (bool, error) {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return false, fmt.Errorf("get connection: %w", err)
	}
	defer conn.Close()
	var has uint64
	q, a := userHasThisLinkQuery(linkToCheck, userID)
	err = conn.QueryRow(q, a...).Scan(&has)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (repo *repository) InsertLink(ctx context.Context, link create.NewLink) error {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("get connection: %w", err)
	}
	defer conn.Close()
	q, a := insertLinkQuery(link.Path, link.RedirectTo, link.UserID)
	_, err = conn.Exec(q, a...)
	return err
}

func (repo *repository) PathExist(ctx context.Context, path string) (bool, error) {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return false, fmt.Errorf("get connection: %w", err)
	}
	defer conn.Close()

	var exist uint64
	q, a := pathExistQuery(path)
	err = conn.QueryRow(q, a...).Scan(&exist)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (repo *repository) UserOwnThisPath(ctx context.Context, path string, userID int64) (bool, error) {
	return false, errors.New("unimplemented")
}

func (repo *repository) DeactivatePath(ctx context.Context, path string) error {
	return errors.New("unimplemented")
}

func (repo *repository) ActivatePath(ctx context.Context, path string) error {
	return errors.New("unimplemented")
}

func (repo *repository) ListLinks(ctx context.Context, option list.SelectOption) ([]list.Link, error) {
	return nil, errors.New("unimplemented")
}
