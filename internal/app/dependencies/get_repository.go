package dependencies

import (
	"link-contractor-api/internal/app/mode/vk"
	"link-contractor-api/internal/dal/pool"
	"link-contractor-api/internal/dal/repository/link"
)

// GetLinkRepository ...
func GetLinkRepository() vk.LinkRepo {
	return link.New(pool.GetPool())
}
