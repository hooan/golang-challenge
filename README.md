# Challenge 

This repository contains a .

## Software needed

[![Git Reference](https://www.svgrepo.com/show/349375/github.svg)](https://github.com)
[![Go Reference](https://www.svgrepo.com/show/349380/go.svg)](https://go.dev/dl/)
[![Docker Reference](https://www.svgrepo.com/show/349342/docker.svg)](https://docker.com)

## Clone the project

```
$ git clone https://github.com/hooan/golang-challenge.git
$ cd golang-challenge
```

## Local 
            
[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

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
