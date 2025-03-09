# Vyking Sample App

## App structure
Main logic is located in /pkg/router/handlers and /pkg/services other things are models, utilities, logger, db connection etc...

Services deal with database by making queries and executing procedures.
I am not validating input data thats something that should be done
when the app gets more complicated and same service function is called from multiple places

Handlers are function that accept the request, parse the payload if there is any, call service function and return the response.

## How to run

First of all you should have golang installed, if you don't, checkout this [link](https://go.dev/doc/install).
Second, you should have docker installed and make sure you can run `docker compose`,
if you don't check out this [link](https://docs.docker.com/engine/install/).

> Note if you don't have make installed check Makefile for commands.

When you have everything set up run these commands (Make sure you are in the project's root folder):
1. Start mysql database by running `make db`, this will start mysql db in background and create a user that matches the DB_DSN in .env
2. Wait couple of seconds for mysql to start and run all `make migrate` to create all tables and stored procedures
3. Run the app by running `make run` or `make`
4. Stop the app by `CTRL+c` and then run `make clean` to stop the docker container and remove the volume

## TODO

- [ ] Add unique constraint on players.email
