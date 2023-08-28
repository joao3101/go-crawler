# go-crawler

The goal of this project is to get all links on the same domain inside a initial provided url.

## Installation

Please change the [config.yaml](./config/config.yaml) with the wanted URL to crawl and the number of simultaneously Go Routines you want to run. 

To install the dependencies.
```bash
go mod tidy
```

To install `pre-commit`.
```bash
pre-commit install
```

Before any commits, always run:
```bash
pre-commit run --all-files
```
This will check for lint being according to standards of the project and also check if go.mod is the way we want it to be.

Other useful commands, such as unit tests, are on the [Makefile](./Makefile).

## Usage

```bash
go run ./cmd/main.go
```

Or, you can run it on Docker. First build the image:
```bash
docker build -t go-crawler .
```

Than run it:
Or, you can run it on Docker. First build the image:
```bash
docker run go-crawler
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)