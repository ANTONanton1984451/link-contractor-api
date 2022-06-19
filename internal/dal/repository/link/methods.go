package link

import (
	"context"
	"errors"
	"fmt"
	"link-contractor-api/internal/entities/user"
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

func (repo *repository) InsertLink(ctx context.Context, link create.NewLink) (err error) {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("get connection: %w", err)
	}
	tx, err := conn.BeginEx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx: %w", err)
	}

	// todo в клозер отдельный в идеале, но можно и повременить пока что
	defer func() {
		var closeErr error

		if err != nil {
			closeErr = tx.Rollback()
		} else {
			closeErr = tx.CommitEx(ctx)
		}

		if closeErr != nil {
			err = closeErr
		}

		conn.Close()
	}()

	q, a := insertLinkQuery(link.Path, link.RedirectTo, link.UserID, link.Type)
	_, err = conn.Exec(q, a...)
	if err != nil {
		// todo так лучше не делать, иначе теряется переиспользуемость данной функции
		// тоже самое и со структурами
		// лучше всего делать приватные обобщённые методы и уже публичные методы, заточенные под юзкейсы
		if isPathExistErr(err) {
			return create.PathIsBusy
		}
		return err
	}

	return repo.redirect.Store(ctx, link.Path, link.RedirectTo)
}

func (repo *repository) UserOwnThisPath(ctx context.Context, path string, userID int64) (bool, error) {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return false, fmt.Errorf("get connection: %w", err)
	}
	defer conn.Close()

	var own uint64
	q, a := userOwnThisPathQuery(path, userID)
	err = conn.QueryRow(q, a...).Scan(&own)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (repo *repository) DeactivatePath(ctx context.Context, path string) error {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("get connection: %w", err)
	}
	defer conn.Close()

	q, a := deactivatePathQuery(path)
	_, err = conn.Exec(q, a...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (repo *repository) ActivatePath(ctx context.Context, path string) error {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("get connection: %w", err)
	}
	defer conn.Close()

	q, a := activatePathQuery(path)
	_, err = conn.Exec(q, a...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (repo *repository) ListLinks(ctx context.Context, usr user.User, option list.SelectOption) ([]list.Link, error) {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return nil, fmt.Errorf("get connection: %w", err)
	}
	defer conn.Close()

	q, a := listLinks(option.OnlyActive, usr.ID)

	rows, err := conn.Query(q, a...)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	linkList := make([]list.Link, 0)
	for rows.Next() {

		var lR linkRow
		err := rows.Scan(&lR.Path, &lR.RedirectTo, &lR.Active, &lR.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("scan rows")
		}

		linkList = append(linkList, mapLinkRowToUcLink(lR))
	}

	return linkList, nil
}
