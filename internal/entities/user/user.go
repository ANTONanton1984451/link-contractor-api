package user

import "time"

type (
	User struct {
		Name         string
		Surname      string
		ID           int64
		ExternalID   string
		SourceType   SourceType
		RegisteredAt time.Time
	}

	SourceType uint64
)

const (
	DevSource SourceType = iota + 1
	VkSource
)
