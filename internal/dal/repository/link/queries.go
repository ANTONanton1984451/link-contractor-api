package link

func insertLinkQuery(path, redirectTo string, personID int64) (string, []interface{}) {
	return `INSERT INTO link(path,redirect_to,owner_id) VALUES ($1, $2, $3)`, []interface{}{path, redirectTo, personID}
}

func pathExistQuery(path string) (string, []interface{}) {
	return `SELECT 1 FROM link WHERE path = $1`, []interface{}{path}
}

func userHasThisLinkQuery(link string, userID int64) (string, []interface{}) {
	return `SELECT 1 FROM link WHERE owner_id = $1 AND redirect_to = $2`, []interface{}{userID, link}
}
