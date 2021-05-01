# UMAP

Umap is a web app for shortening your urls. You paste your long url and it convert it to a shorter one that can be reused at any time.

## Prerequisites
- Golang 1.14+
- Make
- MySQL


## Setup procedure

- Create a .env file using env.example file as template

- Run migration
```shell
make goose-up
```

- Build
```shell
make build
```

## start the app
```shell
make run
```

Go to [http://localhost:8080](https://localhost:8080)


## Author

[Gtindo Dev](https://gtindo.dev)

## Licence

[MIT](https://opensource.org/licenses/MIT)

