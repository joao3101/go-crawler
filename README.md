# go-crawler

The goal of this project is to get all links on the same domain inside a initial provided url.

## Installation

Please change the [config.yaml](./config/config.yaml) with the wanted URL to crawl and the number of simultaneously Go Routines you want to run. 

```bash
go mod tidy
```
To install the dependencies.

```bash
go run ./cmd/main.go
```
To run the app

Other useful commands, such as unit tests, are on the [Makefile](./Makefile).

## Usage

```bash
make run
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)