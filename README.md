[My LinkedIn](https://www.linkedin.com/in/kaiomarx/)

# gofinance

- **Programming Languages**: [Go](https://go.dev/), [Typescript](https://www.typescriptlang.org/)
- **Framework**: [Fiber](https://github.com/gofiber/fiber)
- **Database**: [Postgres](https://hub.docker.com/_/postgres/)
- **ORM**: [Gorm](https://gorm.io/)
- **CLI**: [Air](https://github.com/cosmtrek/air)

## Learn More

This project tends to be a useful website for personal managment of finance data, inspired by a video from [Pirate King](https://www.youtube.com/watch?v=-arxoYcRWeM) but with a touch of my
technical and architectural believes, written in the best/beast GO language.

Going against the community's standard framework (gin), I decided to implement a REST API using the safest and most performant framework of all popular web languages, and luckily that framework is written in Go (huff is not JavaScript)
so I built the application described in the video, but thinking about the pillars of programming in Go, ending the project with a functional application with real cases and an easy-to-understand code, commented and standardized by google's style guides.

despite being inspired by someone else's work, all codes, ideas and standards are my responsibility and effort.

## Back-End Diagram

The process of thinking an application is not easy task, not even the small ones. But the world is made of challenges and as a programmer
what I do best is resolve challenges, with that in mind now you can follow the thinking process from THE VERY FIRST CONCEPT TO THE FINALE DIAGRAMS.

ps: all below are ordered by time, oldest to newest, that means as the application grew I had to adapt and rethink the base foundation to what matted most...

![gofinance diagram](https://user-images.githubusercontent.com/105358332/204063914-ccfd762c-bee5-46bb-9253-5b18ba8c54cf.png)

><strong>you know everything in life evolves, this app isn't differente, by the followed weeks and many interactions with the codebase I had betters ideas on how everything would tie up, look them below:</strong>

![secondConcept](https://user-images.githubusercontent.com/105358332/204090434-b74ca542-c390-4ee8-a198-e9b57958ca41.png)

## Front-End Diagram
![front](https://user-images.githubusercontent.com/105358332/204588680-48190424-93f3-449f-a626-fe3a5e5fdd2c.png)

## Final Results

## Overview

- `config/config.go` - Automatically load .env values from gofinance.env file.
- `database/*` - Database related folder, stores connection function and DB instance.
- `handlers/*` - My "controller" folder, handles all the flux of request and response manipulating Database and Fiber Context.
- `model/*` - Static object models serving as data structure and schema to database tables creation by migrations.
- `routes/*` - Group routing and organize REST endpoints calls.
- `web/*` - Folder of React App built on top of Vitejs.
- `.air.conf` - File to active live reloading of main.go compiled binary in temporary folder excluded on exit.
- `.editorconfig` - Part of VScode extension to organize the files and suit for every editor out there.
- `.compose.yaml` - Composition file used to create database and API container, isolating from outside host affairs.
- `.Dockerfile` - File used to create a API image for docker, making the binaries avaiable to docker compose orchestration.
- `.gofinance.demo.env` - File to save application high sensitive environment data.
- `.main.go` - Entrypoint of the project used to call the database connection, http framework and port listeners.

## Endpoints

- `GET:3000/api` - greetings route for you adventurer to have found this api.
- `POST:3000/api/user` - receives data from JSON request body and creates a new user on database.
- `GET:3000/api/user` - returns all users registered on database at the moment.
- `GET:3000/api/user/:id` - returns specific user that matches the id (type uuid) passed.

## Running Locally API

This application requires golang v1.19+.
This application requires docker engine v20.10.21 and docker compose.
This application requires air live reload installed.

Create a `gofinance.env` file similar to [`demo.env`](https://github.com/KaioMarxDEV/gofinance/blob/main/gofinance.demo.env).

>Just Testing the API
```bash
git clone https://github.com/KaioMarxDEV/gofinance.git
cd gofinance
#change URL string host var to db on ./database/connect.go
docker compose up -d db
docker compose up -d myapp
curl http://localhost:3000/api
```

>You want to dev the backend faster
```bash
git clone https://github.com/KaioMarxDEV/gofinance.git
cd gofinance
docker compose up db -d
#change URL string host var to localhost on ./database/connect.go
air -c .air.conf
```

>You want to dev the frontend faster
```bash
git clone https://github.com/KaioMarxDEV/gofinance.git
cd gofinance/web
yarn && yarn dev
```
