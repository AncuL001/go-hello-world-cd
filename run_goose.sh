set -a
. .env

GOOSE_DRIVER=$DB_DRIVER
GOOSE_DBSTRING="postgresql://$DB_USER:$DB_PASS@$DB_HOST:$DB_PORT/$DB_NAME"

set +a

goose -dir migrations "$@"
