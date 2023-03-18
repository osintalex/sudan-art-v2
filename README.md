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

You can also just do `fmt.Println` to get stuff to stoud in go with this set up.

## Notes

Going to be a lot less work if you can provide urls to look up the images by in the browse bit.
Wonder if it's better to just send the image id though and store all the images on the frontend?
Or in go in the backend I can send over a particular file based off the id? I think that could work.

## To dump the db

`pg_restore -d sudan_art_backup ~/Documents/Sudan\ Art/sudan_art_db_dump.pgsql`

https://alphahydrae.com/2021/02/how-to-export-postgresql-data-to-a-json-file/

COPY (
SELECT json_agg(row_to_json(sudan_art_artwork)) :: text
FROM sudan_art_artwork
) to '<User Home>/sudan_art_database.json';

Pretty neat once I got it working ay.
