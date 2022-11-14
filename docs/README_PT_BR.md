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

   <br />
</p>


Português | [English](../README.md)

## 📖 Introdução:

Plutão é um planeta anão e ele representa a transformação, mudanças e reformas.
E o nosso pluto ajuda o projeto que não tem versionamento de bancos de dados a se organizar e e poder contar com esse recurso muito importante nos dias de hoje.

Obrigado [Laravel](https://github.com/laravel/laravel) nosso projeto de migrações foi inspirado em como o laravel lida com isso.

## 🚀 Características:
- Suporta vários tipos de bancos de dados: MySQL, PostgreSQL(Implementado);
- Pequeno o arquivo executavel tem menos de 20mb de tamanho;
- Suporta Linux, Windows e Mac OS;
- Multi Projetos, você pode ter vários projetos se conectando ao mesmo banco de dados, e compartilhando as migrações;
- Utiliza XML como linguagem de entrada, sendo menos burocratico e aceitando instruções sql bem grandes

## 🧰 Como instalar

### Baixe o pacote de instalação mais recente
```bash
wget https://github.com/astrolink/pluto/releases/download/v1.0.7/pluto
mv pluto /usr/local/bin/pluto
```

## 🏃 Utilizando
**Você deve estar na raiz do projeto**

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

## 🐍 Banco de dados

No banco de dados será criado uma tabela **pluto_logs** onde pode ser verificado o que foi ou não rodado e qual é a origem daquela migração

## 🖊️ A Fazer
- [ ] PostgreSQL
- [ ] Melhorar tratamento de erros
- [ ] Cobertura de 100% do código com teste
- [ ] Mais comandos utéis de checagem de saúde

## 📄 Licença

O código fonte em `pluto` está disponível sob o [MIT License](/LICENSE.md).