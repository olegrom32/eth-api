# Tech Assignment

This solution implements Domain-Driven Design and Hexagonal Architecture.

## Directory structure

### Hexagonal/DDD
* `application` - application layer (handlers and commands)
* `domain` - domain layer (model and value objects)
* `infrastructure` - infrastructure layer (repo and adapters)
* `ui` - ui layer (rest controllers)

### Framework
* `config` - application configs, wire configs and providers
* `pkg` - generated contracts and functional testing lib

### Tools
* `bin` - binaries, generated by `make`, must be gitignored in real project.

## Documentation
`swaggo` is used to generate documentation. OpenAPI docs are stored in `docs` directory.

Please use https://editor.swagger.io/ to visualize the docs.

## Running

Remove `test` directory, run `make` to prepare the project.

## Test coverage
`application/handlers/listgroups` is covered with unit tests.

Functional tests are in `test` directory. A custom proprietary testing lib that is able to
mock external APIs, load fixtures and start application is used here.

## Known issue
The task says that `index.percentageChange` is `uint64`. Apparently, this value may be negative (I've seen -45 percent value).