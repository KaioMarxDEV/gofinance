[My LinkedIn](https://www.linkedin.com/in/kaiomarx/)

# gofinance

- **Framework**: [Fiber](https://github.com/gofiber/fiber)
- **Database**: [Postgres (or any relational database)](https://hub.docker.com/_/postgres/)
- **ORM**: [Gorm](https://gorm.io/)
- **CLI**: [Air](https://github.com/cosmtrek/air)

## Learn More

This project tends to be a useful website for personal managment of finance data, inspired by a video from [Pirate King](https://www.youtube.com/watch?v=-arxoYcRWeM&t=455s&ab_channel=PIRATEKING) but with a touch of my
technical and architectural believes, written in the best/beast GO language.

Going against the community's standard framework (gin), I decided to implement a REST API using the safest and most performant framework of all popular web languages, and luckily that framework is written in Go (huff is not JavaScript) 
so I built the application described in the video, but thinking about the pillars of programming in Go, ending the project with a functional application with real cases and an easy-to-understand code, commented and standardized by google's style guides.

## Overview

- `config/config.go` - Automatically load .env values from gofinance.env file.
- `database/*` - Database related folder, stores connection function and DB instance.
- `handlers/*` - My "controller" folder, handles all the flux of request and response manipulating Database and Fiber Context.
- `model/*` - Static object models serving as data structure and schema to database tables creation by migrations.
- `routes/*` - Group routing and organize REST endpoints calls.
- `.air.conf` - File to active live reloading of main.go compiled binary in temporary folder excluded on exit.
- `.editorconfig` - Part of VScode extension to organize the files and suit for every editor out there.
- `.compose.yaml` - Composition file used to create database and API container, isolating from outside host affairs.
- `.Dockerfile` - File used to create a API image for docker, making the binaries avaiable to docker compose orchestration.
- `.gofinance.demo.env` - File to save application high sensitive environment data.
- `.main.go` - Entrypoint of the project used to call the database connection, http framework and port listeners.

## Endpoints

- `:3000/api` - greetings route for you adventurer to have found this api.

## Running Locally

This application requires golang v1.19+.
This application requires docker engine v20.10.21 and docker compose.
This application requires air live reload installed.

```bash
git clone https://github.com/KaioMarxDEV/gofinance.git
cd gofinance
docker compose up db -d
air -c .air.conf
```

Create a `gofinance.env` file similar to [`.env.example`](https://github.com/KaioMarxDEV/gofinance/blob/main/gofinance.demo.env).
