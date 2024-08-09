# Ask Me Something (Service In Development)

## Run Service

- To run the service you will have to add a .env file to the root 
of your project with the following entries:
```
DATABASE_PORT=5432
DATABASE_NAME=askme
DATABASE_USER="{Your Name}"
DATABASE_PASSWORD="{Your Password}"
DATABASE_HOST=localhost
PGADMIN_PORT=8081
PGADMIN_DEFAULT_EMAIL=admin@admin.com
PGADMIN_DEFAULT_PASSWORD="{Your Password}"
```

- Then you need execute:

```
$ go mod tidy
$ docker compose up -d
```

- In the browser:
```
localhost:8081
```

- A pgAdmin page will open and you will have to enter the email 
and password that you defined in the .env.

- After entering the pgAdmin dasboard you will have to configure 
the server:
  - In General it can be the chosen name, for example: askme_docker_db;
  - In Connection:
    - Host Name: db
    - Port: 5432
    - Maintenance database: askme
    - Username: DATABASE_USER="{Your Name}"
    - Password: DATABASE_PASSWORD="{Your Password}"


- In the project directory:
```
$ go generate 
```

- After that you can query the database interface and you can see that 
the tables have already been created.

- If you prefer to use the terminal to access the database you have to 
- install PostgreSQL and then:

```
$ psql -h localhost -p 5432 -U <DATABASE_USER> -d <DATABASE_NAME>
```