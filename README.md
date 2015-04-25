# icarus-backend
Go Backend for Icarus

## Installation

**Do not clone with git!**

Please run

```bash
go get github.com/Pitt-CSC/icarus-backend
```

to download the files, which will install them into the correct location (`$GOPATH/src/github.com/Pitt-CSC/icarus-backend`)

### Dependencies

Despite having a Goopfile, we're currently not actually using Goop.  The dependencies thus far are:

- [go-socket.io](https://github.com/googollee/go-socket.io)
- [gorm](https://github.com/jinzhu/gorm)

Please install them both with `go get` after cloning the repo.

### Setup

Before running any code, you need to migrate the database.  You can do this using

```bash
go run lib/migrate.go
```

You will also need to re-run this script any time a model's structure is updated, and whenever new models are added.

## Running Code

```
go build
./icarus-backend
```
