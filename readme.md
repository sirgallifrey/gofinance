TODO

# dependencies

## Templ

Golang templates

## air

Hot reloading for the Go application

## Atlas

SQL Migration generation tool

## SQLc

Generate code from SQL

## xc

Task runner that forces you to document based on README.md

## docker, docker-compose

Optional dependencies for development only, used to setup the database

# Running tasks

All the tasks necessary for development and running the project are described bellow

Those tasks are compatible with xc, so if you have it installed, it's as easy as running `xc build` or `xc setup-env` for more info visit: [xcfile.dev](https://xcfile.dev/)

## tasks

### setup-env

interactive: true

```bash
    go install github.com/a-h/templ/cmd/templ@latest
    go install github.com/cosmtrek/air@latest
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

    # This will not work for windows, please refer to atlas docs for how to install it on windows
    curl -sSf https://atlasgo.sh | sh
```

### migration-init

```bash
    go run cmd/cli/cli.go -config=.config.local.yaml db init
```

### migration-run

```bash
    go run cmd/cli/cli.go -config=.config.local.yaml db migrate
```

### migration-rollback

```bash
    go run cmd/cli/cli.go -config=.config.local.yaml db rollback
```

### build

```bash
    go build -o .build/requirementor ./cmd/start_web_app.go
```

### start

interactive: true

```bash
    air
```

### gen-view

Inputs: FILE

```bash
    templ generate -f $FILE
```

### gen-views

```bash
    templ generate
```

### gen-repos

```bash
    sqlc generate
```

### db-apply

Only for local development, directly applies schema changes. Make it very easy for development.

interactive: true

```bash
    atlas schema apply --env local
```

### db-migrate

```bash
    atlas migrate --env local
```

### db-migrate-diff

Shows diff in migrations

```bash
    atlas migrate diff --env local
```

### test

```
    # todo
```
