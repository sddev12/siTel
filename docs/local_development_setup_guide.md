# Local Development Setup Guide

This document covers everything you need to set up for local development
Setup for each service is defined below

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

### Setup Redis
The session service leverages Redis to store session IDs, each associated with a TTL (Time To Live).

You can run redis locally with docker

```
docker run --name sitel-redis-dev -d -p 6379:6379 redis/redis-stack:latest
```

### Setup environment variables
The session service uses the dotenv package which allows for environment variables to be declared in a `.env` file. 

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

## todo: 

Write section for sitel-iam