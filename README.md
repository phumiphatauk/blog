# Blog API 📝

A simple Golang Application API for managing blog data that you can easily adapt.

## Features ✨

- 👥 Manage application user data
- 🔐 Manage user permissions
- 📝 Manage blog data of the application
- 🛠️ Developed with SQLC, GIN, Postgres
- 🚀 Use migrate CLI for database migrations
- 📡 Utilize gRPC
- 🌐 Use Gateway for gRPC connection
- 📬 Use Queue with asynq
- 📧 Send Emails
- 📜 Use Swagger and embed Swagger with statik

## Installation 🛠️

Instructions for installing the application

### Prerequisites 📋

- 🐹 Go 1.23.2
- 🐳 Docker
- 🛠️ make

### Installation Steps 📦

1. Clone the project from GitHub:
    ```sh
    git clone https://github.com/phumiphatauk/blog.git
    cd blog
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set environment variables:
    Edit the `.env` file as needed

4. Create a `data-db` folder for storing Postgres data locally:
    ```sh
    mkdir data-db
    ```

5. Install and run Postgres with Docker:
    ```sh
    make postgres
    ```

6. Install Redis with Docker:
    ```sh
    make redis
    ```

7. Run database migrations:
    ```sh
    make migrateup
    ```

8. Run the application:
    ```sh
    make run
    ```

## Usage 🚀

Instructions for using your application

### API Example 🌐

You can access the Swagger UI to test the API at:
```sh
http://localhost:8080/swagger
```

## Testing 🧪

Instructions for running tests

```sh
make test
```