# EasyPoll API

API for the easypoll application

## Getting Started

**DO NOT** edit files in `docs` folder, they are regenerated everytime the schema changes

```bash
# run server
go run cmd/easypoll-server/main.go --scheme http --port 3000

# build server
go build cmd/easypoll-server/main.go
```

## Documentation

Start the server and go to http://localhost:3000/v1/docs

## Generating new documentation

You should always update `swagger.yml` first, then regenerate docs with

```bash
go generate ./restapi
```

## Built With

* [go-swagger](https://goswagger.io/) - Swagger 2.0 implementation for go

## Contributing

Contributions are greatly appreciated. Please fork this repository and open a pull request to add snippets, make grammar tweaks, improve project structure, etc.
