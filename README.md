# icarus-backend
Go Backend for Icarus

## Dependencies

Despite having a Goopfile, we're currently not actually using Goop.  The dependencies thus far are:

- [go-socket.io](https://github.com/googollee/go-socket.io)
- [gorm](https://github.com/jinzhu/gorm)

Please install them both with `go get` after cloning the repo.

## Setup

Before running any code, you need to migrate the database.  You can do this using

```bash
go run lib/migrate.go
```

## Running Code

```
go build
./icarus-backend
```
