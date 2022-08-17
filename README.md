# Pluto

♇ This tool will help you better control a non-versioned database project.

# Usage

```ssh
pluto init
```

Change informations on pluto.yml and migration

To run migrations

```sh
pluto run
```

To run rollback migrations

```sh
pluto rollback step=-1
```

# How to test local

Use golang 1.19+

```sh
go mod tidy
go build
mv pluto /usr/local/bin/pluto
```

On any folder

```sh
pluto init
```