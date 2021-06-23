# go-graphql-example

An example go web server implementing a GraphQL api with a Sqlite database.

![gopher](.github/images/gopher.jpg)

## Usage

Using `go run`:
```shell
> go run main.go
2020/xx/xx xx:xx:xx connect to http://127.0.0.1:8080/ for GraphQL playground
```

Using `go build`:
```shell
> go build -o go-graphql-example .
> ./go-graphql-example
2020/xx/xx xx:xx:xx connect to http://127.0.0.1:8080/ for GraphQL playground
```

You can access the GraphQL playground at: [http://127.0.1:8080/](http://127.0.1:8080/)

[![playground](.github/images/playground.png)](http://127.0.1:8080/)

## Development

After making changes to `ent/schema` or `graph` directory make sure to run the generate command:
```shell
go generate ./...
```

## Built with

- [ent](https://github.com/ent/ent)
- [gqlgen](https://github.com/99designs/gqlgen)
- [go-sqlite3](https://github.com/mattn/go-sqlite3)