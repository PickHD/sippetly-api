# sippetly-api

Simple RESTful API with **GO Framework** (**GIN**)

## Setup
- Clone the project
- Create local **MySQL** databases ex. (sippetly_api_db)
- Create ``.env`` files inside project
- Copy **PORT** and **DB_DSN_URI** from ``sample.env`` to your ``.env``
- Set **PORT** (ex.8080) and set **DB_DSN_URI** (DSN = Data Source Name) for connecting to mysql
  > *See [DSN Configuration](https://github.com/go-sql-driver/mysql#dsn-data-source-name) for references*
- Run command ``make server`` for run the server
