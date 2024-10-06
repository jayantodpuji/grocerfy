## Grocerfy

Grocerfy is a web application to list your to-do grocery.


## Tech Stack
- PostgreSQL
- Golang
- Node v20.15.1 (you need to install same version)
- React (Vite)
- Tailwind
- Daisy UI


## How to run
1. Clone this repo and cd into the repo
2. Create postgresql databases `grocerfy_development` for development and `grocerfy_test` for running test. You can name whatever you want but do not forget to change in the env
3. Run the migration by `go run ./cmd/migration`
4. Run the server by `go run ./cmd/grocerfy`
5. Run the web by `cd web/` and `pnpm run dev`


## APIs
You can check it on the [postman collection](Grocerfy.postman_collection.json)
You need to call Register user, and Login endpoint. Then you can try all the functionality after get the token


## How to test
Run `go clean -testcache && go test ./... -p 1`


## Todo
- [] Dockerize
- [] Add field `price` to todo item
- [] Enable update to list and it's item
- [] Add more test cases in golang code
- [] Enable test in react web app


## Motivation
To practice create golang and react project from scratch
