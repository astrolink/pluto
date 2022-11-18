<p align="center">
   <img src="https://raw.githubusercontent.com/astrolink/pluto/main/docs/images/astrolink-gopher-compress.png" />
</p>

<p align="center">
   <b>Pluto</b> [plutão], Essa ferramenta ajudará você a controlar melhor um projeto sem versionamento de banco de dados.
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

Português | [English](../README.md)

## 📖 Introdução:

Plutão é um planeta anão e ele representa a transformação, mudanças e reformas.
E o nosso pluto ajuda o projeto que não tem versionamento de bancos de dados a se organizar e e poder contar com esse recurso muito importante nos dias de hoje.

Obrigado [Laravel](https://github.com/laravel/laravel) nosso projeto de migrações foi inspirado em como ele lida com isso.

## 🚀 Características:

- Suporta vários tipos de bancos de dados: MySQL e PostgreSQL(Implementado);
- O arquivo executavel tem menos de 20mb de tamanho;
- Suporta Linux, Windows e Mac OS;
- Multi Projetos, você pode ter vários projetos se conectando ao mesmo banco de dados, e compartilhando as migrações;
- Utiliza XML como linguagem de entrada, sendo menos burocratico e aceitando instruções sql bem grandes

## 🧰 Como instalar

Baixe o pacote de instalação mais recente

### Instale no OSX

#### Instale usando o wget

Se você já tem o wget instalado use, caso não tenha baixe ele usando o reposítorio de pacotes do seu S.O, ou utilize o método via CURL

```bash
wget https://github.com/astrolink/pluto/releases/download/v0.1.13/pluto-osx
chmod 755 pluto-osx
mv pluto-osx /usr/local/bin/pluto
```

ou

#### Instale usando o CURL

Se você já tem o CURL instalado use

```bash
curl -L -o pluto https://github.com/astrolink/pluto/releases/download/v0.1.13/pluto-osx
chmod 755 pluto-osx
mv pluto-osx /usr/local/bin/pluto
```

### Instale no linux

Se você já tem o wget instalado use, caso não tenha baixe ele usando o reposítorio de pacotes do seu S.O, ou utilize o método via CURL

```bash
wget https://github.com/astrolink/pluto/releases/download/v0.1.13/pluto-linux
sudo chmod +x pluto-linux
mv pluto-linux /usr/local/bin/pluto
```

ou

#### Install using CURL

Se você já tem o CURL instalado use

```bash
curl -L -o pluto https://github.com/astrolink/pluto/releases/download/v0.1.13/pluto-linux
sudo chmod +x pluto-linux
mv pluto-linux /usr/local/bin/pluto
```

#### Checando a instalação

Verifique se o pluto foi devidamente instalado usando:

```bash
pluto -h
```

## 🏃 Utilizando

**_Você deve estar na raiz do projeto_**

```bash
pluto init
```

Irá criar um arquivo chamado pluto.yml e a pasta migrations com um exemplo

Como será o arquivo criado pelo pluto

```yml
mysql:
  host: "127.0.0.1"
  port: 3306
  database: "api"
  username: "root"
  password: "secret"

log: "mysql"
source: "api" // Qual é a fonte de dados daquela migração
```

Feito isso podemos executar o pluto

```sh
pluto run
```

Com o arquivo pluto.yml configurado com os dados do banco de dados
Ao rodar o run serão executadas todas as migrations que ainda não rodaram

```sh
pluto rollback step=-1
```

Caso seja necessário voltar algum passo você pode rodar o rollback e a instrução contrária será rodada

```sh
pluto make create_users_table
```

Para criar uma nova migration rode o comando make com o nome que a migration vai ter
é sugerido manter o padrão de nomenclatura.

### Como usar

<p align="center">

   <a href="https://github.com/astrolink/pluto">
      <img alt="Releases" src="https://raw.githubusercontent.com/astrolink/pluto/main/docs/images/how-to-use-pluto.gif" />
   </a>
</p>

## 💻 Arquivo de migração

**O arquivo de migração conta com quatro campos**

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
        PLUTO - Criado tabela de usuários
    </description>
</pluto>
```

O arquivo XML deve estar em volta da estrutura pluto esse campo e todos os outros são obrigátorio.

**Database**: é qual conexão ele irá utilizar do pluto.yml (Possíveis mysql e postgre)

**Run**: Instrução SQL que será rodado para o passo a frente

**Rollback**: Instrução SQL que será rodado na reversão

**Description**: Descrição do que é feito naquela migração será salvo no banco de dados

## 🖥️ Banco de dados

No banco de dados será criado uma tabela **pluto_logs** onde pode ser verificado o que foi ou não rodado e qual é a origem daquela migração

## 🔥 Como utilizar o pluto dentro de outro projeto em Go

Primeiro baixe o projeto como um pacote

```bash
go get -v github.com/astrolink/pluto@v0.1.13
go mod vendor
```

Caso queira rodar as migrações a partir de um ponto utilize

```go
package main

import (
	"github.com/astrolink/pluto/general/pluto"
)

func main() {
	pluto.RunMigrations()
}
```

Caso queira rodar o rollback

```go
package main

import (
	"github.com/astrolink/pluto/general/pluto"
)

func main() {
	pluto.RunRollback()
}
```

Dessa forma você pode iniciar o go de dentro do seu projeto e empacotar ele junto.

Lembrando que mesmo nesse modo você, precisa ter a pasta de migrations e o arquivo pluto.yaml

## 🏠 Como testar local

Faça o fork ou baixe o código fonte

Utilize Golang 1.19+

e rode os seguintes comandos

```bash
go mod tidy

// osx
go build -x
file pluto
mv pluto pluto-osx

// linux
GOOS=linux GOARCH=amd64 go build -x
file pluto
mv pluto pluto-linux

// instalar no OSX para teste
mv pluto pluto-osx /usr/local/bin/pluto
```

## ⚡ Como gerar um novo release

O projeto utiliza o git flow, então faça o seguinte:

```bash
git flow release start v0.1.x
git flow release finish 'v0.1.x'
git push origin v0.1.x
```

## 🖊️ A Fazer

- [ ] PostgreSQL
- [ ] Melhorar tratamento de erros
- [ ] Cobertura de 100% do código com teste
- [ ] Mais comandos utéis de checagem de saúde

## 👋 Contribuidores

Muito obrigado a todas que ajudaram:

[![Contribuidores](http://contributors.nn.ci/api?repo=astrolink/pluto)](https://github.com/astrolink/pluto/graphs/contributors)

## 📄 Licença

O código fonte em `pluto` está disponível sob o [MIT License](/LICENSE.md).
