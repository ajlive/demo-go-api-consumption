# demo-go-api-consumption

Extremely simple demo of Go API consumption

- [demo-go-api-consumption](#demo-go-api-consumption)
  - [Running the example](#running-the-example)
  - [Building](#building)
  - [Formatting](#formatting)

## Running the example

In one window, type:

```
go run server/main.go
```

In another, type:

```
go run client/main.go
```

## Building

Note that running with `go run` yields startup times comparable to Python. To build, run:

```
go build -o build/runserver server/main.go && go build -o build/runclient client/main.go
```

Then run the server with:

```
./build/runserver
```

and the client with:

```
./build/runclient
```

## Formatting

Run

```
npm install
npx prettier --use-tabs --no-semi -w .
```
