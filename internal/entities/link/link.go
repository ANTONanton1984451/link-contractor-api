package link

import "time"

// Link бизнес-сущность ссылки
type Link struct {
	Path       string
	RedirectTo string
	Active     bool
	CreatedAt  time.Time
}
