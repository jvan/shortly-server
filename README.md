# short.ly url-shortener API

## installation

To install the go dependencies first set the `GOPATH`

```sh
export GOPATH=$PWD/go
```

Then install the required dependencies

```sh
go get go get github.com/labstack/echo
```

NOTE: The project source code is stored outside the `go/` directory. For the
project to compile you will need to create a symlink inside the `go/src/`
that points to the project source.

```sh
cd go/src/
ln -s ../../shortly-server
ln -s ../../shortly
```


## building and running

From the `go/bin/` directory, simply install the `shortly-server` program.

```sh
go install shortly-server
```

Then simply run the following to start the server

```sh
./shortly-server
```

NOTE: The server listens on port `1323`.
