# Personal notes

## CSV dataset

The CSV dataset was taken from [this](https://gist.github.com/armgilles/194bcff35001e7eb53a2a8b441e8b2c6) GitHub Gist.

## First deliverable

In the first deliverable I have implemented a simple Web API that reads from [pokemon.csv](/pokemon.csv) and exposes the following endpoints:

1. `/pokemon/{id}`: Returns a pokemon's data by id, if exists
2. `/pokemon`: Returns all pokemons' data

I've also implemented a proposal of the Clean Architecture by dividing code in the following packages:

- [`entity`](/entity): Entities layer
- [`usecase`](/usecase): Use cases layer
- [`adapter`](/adapter): Interface adapters layer
- [`data`](/data) and [`httpapi`](/httpapi): Framework and drivers layer

For the HTTP framework, I decided to use [Gin](https://github.com/gin-gonic/gin), mainly because of its popularity and ease of use.

## Second deliverable

In the second deliverable I have:

1. Implemented a `/comments/{id}` endpoint that fetches data from the [`/comments/{id}`](https://jsonplaceholder.typicode.com/comments/1) endpoint of the [JSON Placeholder API](https://jsonplaceholder.typicode.com/). Whenever a comment is fetched, it is saved in the local [`comment.csv`](/comment.csv) file. The next time a user requests that particular comment, it will be loaded from the CSV instead of the JSON Placeholder API, for faster response times.
2. Created a custom HTTP client for making simple GET requests to any URL.
3. Wrote more documentation for functions and methods.
4. Created test cases for the `usecase`, `adapter` and `entity` layers. I'm using [`testify`](https://github.com/stretchr/testify) as testing toolkit.
5. Refactored the codebase:
   - Separated controller logic and route logic into their own packages.
   - Created a router group to have API versioning. Currently the `/v1` is supported.
   - Moved configuration logic to its own package in [`config`](/config).

## Final deliverable

In the final deliverable I have:

1. Implemented a `/pokemons/type/{parity}/items/{items}/items_per_worker/{items_per_worker}` endpoint that returns a JSON array of pokemons filtered by the parity (even or odd) of their ID. This endpoint uses the _worker pool pattern_ to filter the pokemons.

   The `{items}` parameter specifies how many pokemons must at least appear in the array, although if `{items}` is greater than the amount of total even or odd pokemons, all filtered pokemons will be returned.

   The `{items_per_worker}` parameter specifies at most how many pokemons each worker will process.

   The amount of workers that are spawned is computed with the following formula:

   ```
   #workers = ceil(items / items_per_worker)
   ```

   The endpoint is also limited to spawn up to 20 workers. If the amount of workers computed is greater than 20, then the endpoint will return a 400 HTTP error suggesting that the client should decrease the amount of `{items}` or increase the amount of `{items_per_worker}`.

2. Wrote documentation for the new logic.
3. Created test cases for the new logic.

---

Everything below this line is the original README.

# On Demand Academy Go Bootcamp

## Introduction

Thank you for participating in the Golang Bootcamp course!
Here, you'll find instructions for completing your certification.

## The Challenge

The purpose of the challenge is for you to demonstrate your Golang skills. This is your chance to show off everything you've learned during the course!!

You will build and deliver a whole Golang project on your own. We don't want to limit you by providing some fill-in-the-blanks exercises, but instead request you to build it from scratch.
We hope you find this exercise challenging and engaging.

The goal is to build a REST API which must include:

- An endpoint for reading from an external API
  - Write the information in a CSV file
- An endpoint for reading the CSV
  - Display the information as a JSON
- An endpoint for reading the CSV concurrently with some criteria (details below)
- Unit testing for the principal logic
- Follow conventions, best practices
- Clean architecture
- Go routines usage

## Requirements

These are the main requirements we will evaluate:

- Use all that you've learned in the course:
  - Best practices
  - Go basics
  - HTTP handlers
  - Error handling
  - Structs and interfaces
  - Clean architecture
  - Unit testing
  - CSV file fetching
  - Concurrency

## Getting Started

To get started, follow these steps:

1. Fork this project
1. Commit periodically
1. Apply changes according to the reviewer's comments
1. Have fun!

## Deliverables

> If you are interested in filing the Go Bootcamp in your WizelineOS profile, please contact Academy to assign you a mentor who can review you comply with your capstone project.

We provide the delivery dates so you can plan accordingly; please take this challenge seriously and try to make progress constantly.

For the final deliverable, we will provide some feedback. If you are struggling with something, contact the mentors and peers to get help on time. Feel free to use the slack channel available.

## First Deliverable

Based on the self-study material and mentorship covered until this deliverable, we suggest you perform the following:

- Create an API
- Add an endpoint to read from a CSV file
- The CSV should have any information, for example:

```txt
1,bulbasaur
2,ivysaur
3,venusaur
```

- The items in the CSV must have an ID element (int value)
- The endpoint should get information from the CSV by some field **_(example: ID)_**
- The result should be displayed as a response
- Clean architecture proposal
- Use best practices
- Handle the Errors **_(CSV not valid, error connection, etc)_**

> Note: what’s listed in this deliverable is just for guidance and to help you distribute your workload; you can deliver more or fewer items if necessary. However, if you deliver fewer items at this point, you have to cover the remaining tasks in the next deliverable.

## Second Deliverable

Based on the self-study material and mentorship covered until this deliverable, we suggest you perform the following:

- Create a client to consume an external API
- Add an endpoint to consume the external API client
- The information obtained should be stored in the CSV file
- Add unit testing
- Update the endpoint made in the first deliverable to display the result as a JSON
- Refator if needed

> Note: what’s listed in this deliverable is just for guidance and to help you distribute your workload; you can deliver more or fewer items if necessary. However, if you deliver fewer items at this point, you have to cover the remaining tasks in the next deliverable.

## Final Deliverable

- Add a new endpoint
- The endpoint must read items from the CSV concurrently using a worker pool
- The endpoint must support the following query params:

```text
type: Only support "odd" or "even"
items: Is an Int and is the amount of valid items you need to display as a response
items_per_workers: Is an Int and is the amount of valid items the worker should append to the response
```

- Reject the values according to the query param **_type_** (you could use an ID column)
- Instruct the workers to shut down according to the query param **_items_per_workers_** collected
- The result should be displayed as a response
- The response should be displayed when:

  - The workers reached the limit
  - EOF
  - Valid items completed

> Important: this is the final deliverable, so all the requirements must be included.

## Submitting the deliverables

For submitting your work, you should follow these steps:

1. Create a pull request with your code, targeting the master branch of your fork.
2. Fill this [form](https://forms.gle/K3umfgwQj2HQ6pMe9) including the PR’s url
3. Stay tune for feedback
4. Do the changes according to the reviewer's comments

## Documentation

### Must to learn

- [Go Tour](https://tour.golang.org/welcome/1)
- [Go basics](https://www.youtube.com/watch?v=C8LgvuEBraI)
- [Git](https://www.youtube.com/watch?v=USjZcfj8yxE)
- [Tool to practice Git online](https://learngitbranching.js.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [How to write code](https://golang.org/doc/code.html)
- [Go by example](https://gobyexample.com/)
- [Go cheatsheet](http://cht.sh/go/:learn)
- [Any talk by Rob Pike](https://www.youtube.com/results?search_query=rob+pike)
- [The Go Playground](https://play.golang.org/)

### Self-Study Material

- [Golang Docs](https://golang.org/doc/)
- [Constants](https://www.youtube.com/watch?v=lHJ33KvdyN4)
- [Variables](https://www.youtube.com/watch?v=sZoRSbokUE8)
- [Types](https://www.youtube.com/watch?v=pM0-CMysa_M)
- [For Loops](https://www.youtube.com/watch?v=0A5fReZUdRk)
- [Conditional statements: If](https://www.youtube.com/watch?v=QgBYnz6I7p4)
- [Multiple options conditional: Switch](https://www.youtube.com/watch?v=hx9iHend6jM)
- [Arrays and Slices](https://www.youtube.com/watch?v=d_J9jeIUWmI)
- [Clean Architecture](https://medium.com/@manakuro/clean-architecture-with-go-bce409427d31)
- [Maps](https://www.youtube.com/watch?v=p4LS3UdgJA4)
- [Functions](https://www.youtube.com/watch?v=feU9DQNoKGE)
- [Error Handling](https://www.youtube.com/watch?v=26ahsUf4sF8)
- [Structures](https://www.youtube.com/watch?v=w7LzQyvriog)
- [Structs and Functions](https://www.youtube.com/watch?v=RUQADmZdG74)
- [Pointers](https://tour.golang.org/moretypes/1)
- [Methods](https://www.youtube.com/watch?v=nYWa5ECYsTQ)
- [Interfaces](https://tour.golang.org/methods/9)
- [Interfaces](https://gobyexample.com/interfaces)
- [Packages](https://www.youtube.com/watch?v=sf7f4QGkwfE)
- [Failed requests handling](http://www.metabates.com/2015/10/15/handling-http-request-errors-in-go/)
- [Modules](https://www.youtube.com/watch?v=Z1VhG7cf83M)
  - [Part 1 and 2](https://blog.golang.org/using-go-modules)
- [Unit testing](https://golang.org/pkg/testing/)
- [Go tools](https://dominik.honnef.co/posts/2014/12/an_incomplete_list_of_go_tools/)
- [More Go tools](https://dev.to/plutov/go-tools-are-awesome-bom)
- [Functions as values](https://tour.golang.org/moretypes/24)
- [Concurrency (goroutines, channels, workers)](https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3)
  - [Concurrency Part 2](https://www.youtube.com/watch?v=LvgVSSpwND8)
