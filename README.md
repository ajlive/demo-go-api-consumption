# demo-go-api-consumption

Extremely simple demo of Go API consumption

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
go build -o runserver server/main.go && go build -o runclient client/main.go
```

Then run the server with:

```
./runserver
```

and the client with:

```
./runclient
```

## Formatting

Run

```
npm install
npx prettier --use-tabs --no-semi -w .
```
