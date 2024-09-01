# Real time Forum

## Description

Real-time web forum that allows users to register, create posts, comment on posts, and send private chat messages to one another using Gorilla websocket

## Authors

- Ingvar-V. Leerimaa (IngvarLeerimaa)
- Margus Toompere (MargusT)

## Technologies used

Frontend:
- HTML, CSS, JS

Backend:
- Go 1.20
- SQLite3
- Gorilla websocket

## Required modules

- github.com/gorilla/websocket v1.5.0
- github.com/gofrs/uuid v4.4.0+incompatible
- golang.org/x/crypto v0.13.0
- github.com/mattn/go-sqlite3

## Running 

Use genCert.bash to generate crt and key
    
```bash genCert.bash```

Run the server:

`go run main.go`

You will be asked to enter the port number for the server to run on.

# Audit

[Audit points can be found here](https://01.alem.school/git/root/public/src/branch/master/subjects/real-time-forum/audit)

*As time is of the essence, implementation of the project is not polished but has successfully addressed all audit points.*

## Existing users

- Username: asd
- Password: asdasdasd

- Username: qwe
- Password: qweqweqwe

