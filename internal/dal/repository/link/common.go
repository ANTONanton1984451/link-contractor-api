package link

import (
	"link-contractor-api/internal/usecase/link/list"
	"strings"
	"time"
)

type (
	linkRow struct {
		Path       string
		RedirectTo string
		Active     bool
		CreatedAt  time.Time
	}
)

func mapLinkRowToUcLink(lr linkRow) list.Link {
	return list.Link{
		Path:       lr.Path,
		RedirectTo: lr.RedirectTo,
		Active:     lr.Active,
		CreatedAt:  lr.CreatedAt,
	}
}

func isPathExistErr(err error) bool {
	return strings.Contains(err.Error(), _uniquePathConstraint)
}

const (
	_uniquePathConstraint = `"unique_path"`
)
