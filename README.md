## Requirements

- [docker compose](https://docs.docker.com/compose/install/)
- [goose - migration tool](https://github.com/pressly/goose)
- [jet - sql generator](https://github.com/go-jet/jet)

## Getting started

- ### Install goose

```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

*This will install the `goose` binary to your `$GOPATH/bin` directory.*

For macOS users `goose` is available as a [Homebrew Formulae](https://formulae.brew.sh/formula/goose#default):

```shell
brew install goose
```

- ### Install jet

Jet generator can be installed by following way:

```sh
go install github.com/go-jet/jet/v2/cmd/jet@latest
```

*Make sure `dir_path` folder is added to the PATH environment variable.*

- ### Add environment `.env` file to project directory. See `.env.example`.

- ### Run docker compose. Linux example:

```bash
cd path/to/project && docker compose up -d
```

- ### Run makefile from command line. Linux example:

```bash
cd path/to/project && make run
```