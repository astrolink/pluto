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
pluto rollback
```

# How to test local

```sh
go build
mv pluto /usr/local/bin/pluto
```

On any folder

```sh
pluto init
```