## Built with

* [gorm](https://gorm.io/docs/index.html) - ORM library for Golang
* [golang-jwt](https://github.com/golang-jwt/jwt) - Library for working with jwt
* [go-playground/validator](https://github.com/go-playground/validator) - Library for validation
* [go-faker](https://github.com/go-faker/faker) - Library for generating fake data
* [testify](https://github.com/stretchr/testify) - Go code (golang) set of packages that provide many tools for testifying that your code will behave as you intend

## Start application

1. Start the application using Docker Compose:

    ```sh
    docker-compose build
    docker-compose up
    ```

## Full docker rebuild
1) go mod tidy
2) go mod vendor
3) docker-compose down
4) docker-compose build --no-cache
5) docker-compose up