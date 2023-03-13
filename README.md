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
