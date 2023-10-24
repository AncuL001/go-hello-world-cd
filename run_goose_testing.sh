set -a
. .env

GOOSE_DRIVER=postgres
GOOSE_DBSTRING="postgresql://$TESTING_DB_USER:$TESTING_DB_PASS@$TESTING_DB_HOST:$TESTING_DB_PORT/$TESTING_DB_NAME"

set +a

goose -dir migrations "$@"
