# Dating app

This is a simple Dating app service written on Golang and set up to run with Docker Compose. It provides an easy way to develop and deploy your Golang application in a containerized environment.

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Golang](https://golang.org/doc/install) (if you want to build the Golang binary locally)

## Usage

1. Clone the repository:
   > git clone https://github.com/Farhanramadhana/dating-app.git
   > cd dating-app
2. copy .env.example values to .env
3. RUN docker compose up -d --build
   The application should now be accessible at http://localhost:port
4. Execute migration.sql to create database table

##Run Without Docker
If you want to work on the application locally without using Docker, you can run the golang command directly:
go run main.go
