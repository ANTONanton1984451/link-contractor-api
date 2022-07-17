package link

import (
	"context"
	"errors"
	"fmt"
	"link-contractor-api/internal/entities/link"
	"link-contractor-api/internal/entities/user"
	"link-contractor-api/internal/usecase/link/create"
	"link-contractor-api/internal/usecase/link/list"

	"github.com/jackc/pgx"
)

// UserHasThisLink импелементация интерйейса юзкейса создания ссылки
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

// InsertLink импелентация интерфейса создания ссылки
func (repo *repository) InsertLink(ctx context.Context, link create.NewLink) (err error) {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return fmt.Errorf("get connection: %w", err)
	}
	tx, err := conn.BeginEx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx: %w", err)
	}

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
		// проверка, не упёрлись ли мы в ограничения базы на уникальность пути
		if isPathExistErr(err) {
			return create.PathIsBusy
		}
		return err
	}

	return nil
}

// UserOwnThisPath импелментация интерфейса юзкейсов активации и деаквтиации пути
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

// ListLinks импелментация интерфейса для юзкейса получения ссылок
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

func (repo *repository) Get(ctx context.Context, path string) (*link.Link, error) {
	conn, err := repo.pool.GetConn(ctx)
	if err != nil {
		return nil, fmt.Errorf("get connection: %w", err)
	}
	defer conn.Close()

	q, a := getLinkQuery(path)

	var lnk link.Link

	if err = conn.QueryRow(q, a...).Scan(&lnk.Path, &lnk.RedirectTo, &lnk.Active, &lnk.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get query: %w", err)
	}

	return &lnk, nil
}
