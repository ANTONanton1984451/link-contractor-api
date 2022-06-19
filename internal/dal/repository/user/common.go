package user

import (
	userEntity "link-contractor-api/internal/entities/user"
	"time"
)

type (
	userRow struct {
		ID         int64
		Name       string
		Surname    string
		ExternalID string
		CreatedAt  time.Time
		SourceID   uint64
	}
)

const (
	_devSource = `dev`
	_vkSource  = `vk`
)

const (
	_name    = `Anton`
	_surname = `Shushpanov`
)

var entrypointInnerSource = map[userEntity.SourceType]string{
	userEntity.DevSource: _devSource,
	userEntity.VkSource:  _vkSource,
}
