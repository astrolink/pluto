<p align="center">
   <img src="https://raw.githubusercontent.com/astrolink/pluto/main/docs/images/astrolink-gopher-compress.png" />
</p>

<p align="center">
   <b>Pluto</b> ♇ This tool will help you better control a non-versioned database project.
</p>

<p align="center">

   <a href="https://github.com/astrolink/pluto/releases">
      <img alt="Releases" src="https://img.shields.io/github/release/astrolink/pluto.svg?style=flat-square&include_prereleases" />
   </a>

   <a href="https://github.com/astrolink/pluto/releases">
      <img alt="All Releases" src="https://img.shields.io/github/downloads/astrolink/pluto/total.svg?style=flat-square" />
   </a>

   <a href="https://github.com/astrolink/pluto/commits">
      <img alt="Last commit" src="https://img.shields.io/github/last-commit/astrolink/pluto.svg?style=flat-square" />
   </a>

   <img src="https://img.shields.io/github/go-mod/go-version/astrolink/pluto?style=flat-square">

   <a href="https://goreportcard.com/report/github.com/astrolink/pluto">
      <img alt="Go Report" src="https://goreportcard.com/badge/github.com/astrolink/pluto" />
   </a>

   <br />
</p>

English | [Português](docs/README_PT_BR.md)

## 📖 Introduction:

Pluto is a dwarf planet and it represents transformation, change and reform.
And our pluto helps the project that doesn't have database versioning to organize itself and be able to count on this very important resource these days.

Thanks [Laravel](https://github.com/laravel/laravel) our migrations project was inspired by how it handles this.

## 🚀 Features:

- Supports several types of databases: MySQL and PostgreSQL(Implemented);
- The executable file is less than 20mb in size;
- Supports Linux, Windows and Mac OS;
- Multi Projects, you can have several projects connecting to the same database, and sharing migrations;
- Uses XML as input language, being less bureaucratic and accepting very large sql statements

## 🧰 How to install

Download the latest installation package

### Install using wget

If you have wget utilize this

```bash
wget https://github.com/astrolink/pluto/releases/download/v0.1.9/pluto
chmod 755 pluto
mv pluto /usr/local/bin/pluto
```

### Install using CURL

If you have CURL utilize this

```bash
curl -L -o pluto https://github.com/astrolink/pluto/releases/download/v0.1.9/pluto
chmod 755 pluto
mv pluto /usr/local/bin/pluto
```

To check pluto is properly installed use

```bash
pluto -h
```

## 🏃 Using

**_You must be at the root of the project_**

```bash
pluto init
```

It will create a file called pluto.yml and the migrations folder with an example

What the file created by pluto will look like:

```yml
mysql:
  host: "127.0.0.1"
  port: 3306
  database: "api"
  username: "root"
  password: "secret"

log: "mysql"
source: "api" // What is the data source for that migration
```

After that we can run pluto

```sh
pluto run
```

With the pluto.yml file configured with the database data
When running run, all migrations that have not yet run will be executed

```sh
pluto rollback step=-1
```

If it is necessary to go back some step, you can run the rollback and the opposite instruction will be run.

```sh
pluto make create_users_table
```

To create a new migration run the make command with the name that the migration will have
it is suggested to keep the naming standard.

### How to use

<p align="center">

   <a href="https://github.com/astrolink/pluto">
      <img alt="Releases" src="https://raw.githubusercontent.com/astrolink/pluto/main/docs/images/how-to-use-pluto.gif" />
   </a>
</p>

## 💻 Migration file

**The migration file has four fields**

```xml
<?xml version="1.0" encoding="UTF-8"?>
<pluto>
    <database>
        mysql
    </database>
    <run>
        CREATE TABLE users (name VARCHAR(20),email VARCHAR(20),created_at DATE);
    </run>
    <rollback>
        DROP TABLE users;
    </rollback>
    <description>
        PLUTO - User table created
    </description>
</pluto>
```

The XML file must be around the pluto structure this field and all others are mandatory.

**Database**: is which connection it will use from pluto.yml (Possible mysql and postgre)

**Run**: SQL statement that will be run for the next step

**Rollback**: SQL statement that will be run on rollback

**Description**: Description of what is done in that migration will be saved in the database

## 🖥️ Database

A **pluto_logs** table will be created in the database where you can check what was or was not run and what is the origin of that migration

## 🔥 How to use pluto inside another project in Go

First download the project as a package

```bash
go get -v github.com/astrolink/pluto@v0.1.9
go mod vendor
```

If you want to run the migrations from a point, use

```go
package main

import (
	"github.com/astrolink/pluto/general/pluto"
)

func main() {
	pluto.RunMigrations()
}
```

If you want to run the rollback

```go
package main

import (
	"github.com/astrolink/pluto/general/pluto"
)

func main() {
	pluto.RunRollback()
}
```

That way you can launch go from within your project and package it together.

Remembering that even in this mode you need to have the migrations folder and the pluto.yaml file

## 🏠 How to test location

Fork or download the source code

Use Golang 1.19+

and run the following commands

```bash
go mod tidy
go build
mv pluto /usr/local/bin/pluto
```

## ⚡ How to generate a new release

The project uses git flow, so do the following:

```bash
git flow release start v0.1.x
git flow release finish 'v0.1.x'
git push origin v0.1.x
```

## 🖊️ To Do

- [ ] PostgreSQL
- [ ] Improve error handling
- [ ] 100% code coverage with testing
- [ ] More useful health check commands

## 👋 Contributors

Many thanks to everyone who helped:

[![Contributors](http://contributors.nn.ci/api?repo=astrolink/pluto)](https://github.com/astrolink/pluto/graphs/contributors)

## 📄 License

The source code in `pluto` is available under the [MIT License](/LICENSE.md).
