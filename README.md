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

## Notes

Need to next validate image name with regex then check it can return stuff.
Need to do a general refactor of this now it's working...
Things to change:

- general style and logging pass - should use log not fmt everywhere, check that gives output
- add all the b64 data to the json, that's a much neater pattern. Code for that in a utils script, also then
remove the json from version control.
- Then need to... bump up this issue https://github.com/netlify/netlify-lambda/issues/630
- Then once that's all done, need to add a get imagedetails endpoint for when you click on the image...
or do you? that's all in the json already so could just get set in the frontend really
- if you do that then just throw error if url path is wrong
## To dump the db

`pg_restore -d sudan_art_backup ~/Documents/Sudan\ Art/sudan_art_db_dump.pgsql`

https://alphahydrae.com/2021/02/how-to-export-postgresql-data-to-a-json-file/

COPY (
SELECT json_agg(row_to_json(sudan_art_artwork)) :: text
FROM sudan_art_artwork
) to '<User Home>/sudan_art_database.json';

Pretty neat once I got it working ay.
