query="CREATE TABLE IF NOT EXISTS users ( \
  acct VARCHAR(50) PRIMARY KEY, \
  pwd TEXT NOT NULL, \
  fullname VARCHAR(50) NOT NULL, \
  created_at TIMESTAMP, \
  updated_at TIMESTAMP \
);"

if [ -z "$1" ]
then
    echo "first argument must be password"
    exit 1
fi

export PGPASSWORD=$1;

psql -h localhost -U test -d test -c "$query"