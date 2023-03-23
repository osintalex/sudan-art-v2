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

## Notes

- Then need to... bump up this issue https://github.com/netlify/netlify-lambda/issues/630
- Write docstrings based off this https://github.com/spf13/cobra/blob/main/args.go
- Run some go lint checks
- Deploy
-Change repo name again and update medium article

## To dump the db

`pg_restore -d sudan_art_backup ~/Documents/Sudan\ Art/sudan_art_db_dump.pgsql`

https://alphahydrae.com/2021/02/how-to-export-postgresql-data-to-a-json-file/

COPY (
SELECT json_agg(row_to_json(sudan_art_artwork)) :: text
FROM sudan_art_artwork
) to '<User Home>/sudan_art_database.json';

Pretty neat once I got it working ay.
