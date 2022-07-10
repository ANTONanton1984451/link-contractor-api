echo "Start migrations"

env $(cat config/prod.env | xargs) bash scripts/migrations/migrate.sh

echo "Start app"
bin/main --config_path $1
