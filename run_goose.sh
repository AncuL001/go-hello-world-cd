set -a
. .env

GOOSE_DRIVER=postgres
GOOSE_DBSTRING="postgresql://$DB_USER:$DB_PASS@$DB_HOST:$DB_PORT/$DB_NAME"

set +a

goose -dir migrations "$@"
