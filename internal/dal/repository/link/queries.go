package link

import "link-contractor-api/internal/usecase/link/create"

const (
	_randomType        = `random`
	_userGeneratedType = `user_generated`
)

func insertLinkQuery(path, redirectTo string, personID int64, lType create.LinkType) (string, []interface{}) {
	dbType := _randomType
	switch lType {
	case create.UserGenerated:
		dbType = _userGeneratedType
	}
	return `INSERT INTO link(path,redirect_to,owner_id,create_type) VALUES ($1, $2, $3, $4)`, []interface{}{path, redirectTo, personID, dbType}
}

func pathExistQuery(path string) (string, []interface{}) {
	return `SELECT 1 FROM link WHERE path = $1`, []interface{}{path}
}

func userHasThisLinkQuery(link string, userID int64) (string, []interface{}) {
	return `SELECT 1 FROM link WHERE owner_id = $1 AND redirect_to = $2`, []interface{}{userID, link}
}

func userOwnThisPathQuery(path string, userID int64) (string, []interface{}) {
	return `SELECT 1 FROM link WHERE owner_id = $1 AND path = $2`, []interface{}{userID, path}
}

func deactivatePathQuery(path string) (string, []interface{}) {
	return `UPDATE link SET active = false WHERE path = $1`, []interface{}{path}
}

func activatePathQuery(path string) (string, []interface{}) {
	return `UPDATE link SET active = true WHERE path = $1`, []interface{}{path}
}

func listLinks(onlyActive bool, userID int64) (string, []interface{}) {
	if onlyActive {
		return `SELECT path,redirect_to,active,created_at FROM link WHERE active = true AND owner_id = $1`, []interface{}{userID}
	}
	return `SELECT path,redirect_to,active,created_at FROM link WHERE owner_id = $1`, []interface{}{userID}
}

func getLinkQuery(path string) (string, []interface{}) {
	return `SELECT path,redirect_to,active,created_at FROM link WHERE path = $1`, []interface{}{path}
}
