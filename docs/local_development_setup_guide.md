# Local Development Setup Guide

This document covers everything you need to set up for local development

Set up for each service is explained below

## sitel-frontend

Install [NodeJs](https://nodejs.org/en) if not alerady installed

### Install dependecies
```
cd sitel-frontend
npm install
```

### Run app locally
```
npm run dev
```

<br>

## sitel-session
Install [NodeJs](https://nodejs.org/en) if not alerady installed

### Install dependecies
```
cd sitel-frontend
npm install
```

### Set up Redis
The session service leverages Redis to store session IDs, each associated with a TTL (Time To Live)

You can run redis locally with docker

```
docker run --name sitel-redis-dev -d -p 6379:6379 redis/redis-stack:latest
```

### Set up environment variables
The session service uses the [dotenv](https://www.npmjs.com/package/dotenv) package which allows for environment variables to be declared in a `.env` file. 

Create a file in `./sitel-session` with the `REDIS_HOST` and `REDIS_PORT` variables defined
```
cd ./sitel-session
echo "REDIS_HOST=localhost" > .env
echo "REDIS_PORT=6379" >> .env
```

### Run app locally
```
npm run dev
```

<br>

## sitel-iam
Install [Go](https://go.dev/) if not alredy installed

### Install dependencies
```
cd ./sitel-iam
go get download
```

### Setup Mongo
The sitel-iam service leverages MongoDB to store and retrieve user data

You can run Mongo locally with Docker
```
docker run --name sitel-mongo -d -p 27017:27017 mongo:latest
```

### Set up environment variables
The sitel-iam services uses the [godotenv](https://github.com/joho/godotenv) package which allows for environment variables to be declared in a `.env` file

Create a `.env` file in `./sitel-session` with the `MONGO_DB_URI`, `MONGO_DB_PORT` `LOGGING_LEVEL` and `SITEL_SESSION_HOST` variables defined
```
cd ./sitel-session
echo "MONGO_DB_HOST=localhost > .env
echo "MONGO_DB_PORT=27017" >> .env
echo "LOGGING_LEVEL=INFO" >> .env
echo "SITEL_SESSION_HOST=localhost" >> .env
```

### Run app locally
```
go run main.go
```

<br>

## Docker Compose
You can spin up all services locally with docker compose from the root of the repo
```
docker compose up -d --build
```

Tear down
```
docker compose down --remove-orphans
```
