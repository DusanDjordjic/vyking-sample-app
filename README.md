# Vyking Sample App

## App structure
Main logic is located in /pkg/router/handlers and /pkg/services other things are models, utilities, logger, db connection etc...
Queries are located in /pkg/db/queries and are placed into the binary on compile time with go:embed package.

Services deal with database by making queries and executing procedures.
I am not validating input data there, but thats something that should be done
when the app gets more complicated and same service function is called from multiple places

Handlers are function that accept the request, parse the payload if there is any, call service function and return the response.

## How to run

First of all you should have golang installed, if you don't, checkout this [link](https://go.dev/doc/install).
Second, you should have docker installed and make sure you can run `docker compose`,
if you don't check out this [link](https://docs.docker.com/engine/install/).

> Note if you don't have make installed check Makefile for commands.

When you have everything set up run these commands (Make sure you are in the project's root folder):

1. From projects root run `go mod tidy` to install all needed packages
2. Start mysql database by running `make db`, this will start mysql db in background and create a user that matches the DB_DSN in .env
3. Wait couple of seconds for mysql to start and run `make migrate` to create all tables and stored procedures
4. Run the app by running `make run` or `make`
5. Stop the app by `CTRL+c` and then run `make clean` to stop the docker container and remove the volume

## Q&A

### QUESTION: What did you learn and if you encountered any challenges, how did you overcome them?

### ANSWER

I haven't learned much about golang because I already knew almost everything but I haven't used stored procedures much in the past,
I was starting transactions from golang and doing everything in the code. I knew some basic stuff already but this was great a learning experience.
There are many benefints from doing it like this and I see them now.

First thing that's really important is that it's faster because there is no need transfer the whole query over the network and there is no need
for database to parse the query again and again which is great. Saying that it should be probably better to always use stored procedures when you have
static queries with parameters (making sure that db doesn't use to much memory for caching them).
Example of "dynamic" query would be some route that can filter and sort data if user sent some query parameters
(maybe that's also possible from stored procedures, I don't know.).

Second thing is that it adds another level of security and also another layer for devs to handle errors which I like very much. Also it should be
possible to allow users only to interact with a table by calling stored procedures and not by executing some arbitrary queries which will
bring the security to the another level.



### QUESTION: What did you take in consideration to ensure the query and or stored procedure is efficient and handles edge cases?

### ANSWER

I wrote the queries the best I know but for speeding it up adding indexes on columns that are used for filtering and sorting will help.
Looking at the execution plan is crucial at this phase and trying to understand what's going on. There are some obvious speedups, for example
adding an index on (player_id, tournament_id) on tournament_bets table because RANK groups by tournament_id and then we group by player_id and tournament_id.
But before any "non-obvious" changes a real environment should be emulated, to see how fast are we running currently and taking it on from there.

I created the ranking system both by using window functions and CTEs and stored procedures and I'm not really sure what's faster, that will have to be tested.
Why I did that? Because a stored procedure would offer more flexibility in my opinion, for example, adding a feature where for every tournament we have different prize distribution and
we may give prizes to 4th and 5th place for example. Maybe that's possible with plain query but it was easier for me to think about that through stored procedure.

About the edge cases, I have checks that make sure tournament is running, user has enough funds but I haven't solved the case when users share a place.
I guess that could be solved by for example giving the prize to the user who has the last bet or maybe the first one or maybe giving the prize to both of them.
I think that giving both players a prize is better than choosing one of them, because I know how I would feel if I lost my prize like that :D.


### QUESTION: If you used CTEs or Window Functions, what did you learn about their power and flexibility?

### ANSWER

Like with stored procedures I didn't have a need for more complex queries in the past so I knew how to use the Window Functions but
it took some time to create the ranking query. Window function are great but I must say that I don't know most of them
so I cannot really say anything about their power for sure but, my guess is that you can use them for many things.

CTEs are awesome, flexibility and power are great because you can do whatever you want with them.
Every time you need some temporary result you can use them.


### QUESTION: How might you apply the technique in more complex scenarios

### ANSWER

For more complex scenarios I whould do the same I did here, break down the logic that I need to create into smaller steps and take it from there.
First everything that I know how to do I would do right away and then one by one solve the things I didn't know right away.
CTEs are great for having temporary results so I would recognize the places that I need them for and do it.
I my experience writing queries longer than like 5-6 lines is always like this.

### QUESTION: Optimization: How did you optimize the queries and stored procedures.

### ANSWER

Like I said above there are some obvious things that could be done right away,
but apart from those it's hard to optimize something that's running on very small scale because it takes like few micro seconds to do the query,
so first of  all I would emulate a real enviroment. Then I would test it and keep notes about how fast every thing is running.
The ones that are slowest I would optimize first. It's same with code optimizations, there are a lot of things that impact performace so only with
careful testing you can be sure that something is faster.
