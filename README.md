# Go REST API

Sample Go REST API app built with Fiber. JWT auth (without refresh token). GORM with Postgres DB. Apitest for integration testing. Air for hot reload in docker dev mode. @ stage dockerfile for prod mode build.
(_misc folder is just some concepts for go studying)

## Install
- `docker` and `docker compose` should be up and running (it's required for Postgres)
- `make` comand should be available (but it's optionally)

## How to use
- first start `make devup`
- any further starts `make dev`
- stop `make stop`
- stop and delete docker containers `make devdn`

## Check with CURL

```sh
# signup and get jwt from response
curl -XPOST -H "Content-type: application/json" -d '{"email":"test@test.com","password":"123123"}' 'http://localhost:9001/api/v1/signup'

# create a couple of items
curl -XPOST -H 'Authorization: Bearer PASTE_JWT' -H "Content-type: application/json" -d '{"name":"coffee","price":10,"quantity":10}' 'http://localhost:9001/api/v1/items'

curl -XPOST -H 'Authorization: Bearer PASTE_JWT' -H "Content-type: application/json" -d '{"name":"milk","price":299,"quantity":100}' 'http://localhost:9001/api/v1/items'

# check items
curl -XGET -H "Content-type: application/json" 'http://localhost:9001/api/v1/items'
```