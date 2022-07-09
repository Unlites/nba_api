#!/bin/sh
# wait_for_postgres.sh

set -e

host="$1"
pass="$2"
cmd="$3"

until PGPASSWORD=$pass psql -h "$host" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd