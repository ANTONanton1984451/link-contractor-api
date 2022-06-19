package user

import (
	userEntity "link-contractor-api/internal/entities/user"
)

func getUserQuery(externalID string, source string) (string, []interface{}) {
	return `SELECT id,name,surname,source_id,created_at 
			FROM person WHERE external_id = $1 AND source_id = 
			(SELECT id FROM source WHERE name = $2)`,
		[]interface{}{externalID, source}
}

func createUserQuery(user userEntity.User) (string, []interface{}) {
	return `INSERT INTO person(name,surname,external_id,source_id) 
			VALUES($1,$2,$3,(SELECT id FROM source WHERE name = $4)) RETURNING id`,
		[]interface{}{user.Name, user.Surname, user.ExternalID, entrypointInnerSource[user.SourceType]}
}

func getUserFromBanQuery(user userEntity.User) (string, []interface{}) {
	return `SELECT until FROM ban_list WHERE person_id = $1 AND ban_active = true`, []interface{}{user.ID}
}
