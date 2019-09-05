# EasyPoll API

API for the easypoll application

## Getting Started

**DO NOT** edit files in `docs` folder, they are regenerated everytime the schema changes

```bash
# clone project
go get github.com/wunari/easypoll-backend

# run server
go run main.go

# build server
go build
```

## Documentation

Start the server and go to http://localhost:3000/v1/docs

## Generating new documentation

You should always update `swagger.yml` first, then regenerate docs with

```bash
swagger generate server -t docs -f ./docs/swagger.yml --exclude-main -A easypoll
```

## Live version

https://wunari-easypoll.herokuapp.com/v1/docs

## Built With

* [go-swagger](https://goswagger.io/) - Swagger 2.0 implementation for go

## Contributing

Contributions are greatly appreciated. Please fork this repository and open a pull request to add snippets, make grammar tweaks, improve project structure, etc.
