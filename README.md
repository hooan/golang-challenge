# Challenge 

This repository contains a .

## Software needed

[Git](https://github.com)
[Go](https://go.dev/dl/)
[Docker Reference](https://docker.com)

## Clone the project

```
$ git clone https://github.com/hooan/golang-challenge.git
$ cd golang-challenge
```

## Local 
            
 * Configure db (app/config.go)
 * Create an .env file 
      ```
            TRANSACTIONS_FILE="repository/txns.csv"
            TO="some.email@gmail.com"
            EXECUTION_ENV="local"
            FROM="test.other.email@outlook.com"
            HOST="smtp-mail.outlook.com"
            PORT="587"
            PASSWORD="password"
            USERNAME="test.other.email@outlook.com"
      ```
 * Run project

```
$ make run

```

## Docker 
      * Run docker compose
            ```
            $ docker compose up

            ```
