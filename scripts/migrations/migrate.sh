cd scripts/migrations

goose postgres "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=${PGSSLMODE} host=${POSTGRES_HOST}" up
