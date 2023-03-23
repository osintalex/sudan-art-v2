# Set Up

Dependencies:

1. Install [goenv](https://github.com/syndbg/goenv).
2. Install [nvm](https://github.com/nvm-sh/nvm).

```shell
# Install go dependencies for backend
goenv install 1.20.1
goenv local 1.20.1
cd functions/v1
go mod download
cd ../..

# Install node dependencies for frontend
nvm use
npm i

# Now install netlify cli to run this locally
npm install netlify-cli -g
netlify dev
```

## Deployment

Just run `netlify deploy`.

## To dump the db

This project previously had a postgres database as the backend. I got rid of this to reduce costs.

For reference, this his how:

```shell
pg_restore -d sudan_art_backup ~/Documents/Sudan\ Art/sudan_art_db_dump.pgsql
```
```sql
COPY (
SELECT json_agg(row_to_json(sudan_art_artwork)) :: text
FROM sudan_art_artwork
) to './sudan_art_database.json';
```
