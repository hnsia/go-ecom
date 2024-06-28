# GO-ECOM

This is a simple e-commerce project written in GO that contains common coding conventions in GO. It only contains a REST API server without any front-end.

## Installation

Install the following tools to run the project.

- [Migrate (for DB migrations)](https://github.com/golang-migrate/migrate/tree/v4.17.0/cmd/migrate)

## Running the project

1. Install and setup a MySQL database.
1. Create a database with the desired name such as 'go-ecom'
1. Run migrations

```bash
make migrate-up
```

4. Run the project with the following command

```bash
make run
```

## Tests

Run the tests using the following command:

```bash
make test
```
