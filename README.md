# select-location

## Features
- Import a json file that contains locations
- Typeahead suggestion
- Google Map integration

This project utilized Go, vue.js, docker, and GraphQL

## Installation
Please make sure Go, Docker, vue3 is properly installed

### Backend
#### Install dependencies:
```sh
go get github.com/graphql-go/graphql
go get go.mongodb.org/mongo-driver/mongo
go get github.com/rs/cors
```

#### start MongoDB
```sh
docker-compose up -d
```

#### Data seed:
Redirect to ./backend/tools/seed
```sh
go run seed.go
```
#### Verify data seed:
```sh
docker exec -it YOUR_CONTAINER_NAME mongosh
use us_states db.states.find() 
db.states.find()
```
#### Start GraphQL server:
```sh
cd backend/cmd/server
go run main.go
```
#### Run test
```sh
cd backend/
go test ./tests/... -v
```
#### Test API:
```sh
curl 'http://localhost:8080/graphql?query={states{name}}'
curl 'http://localhost:8080/graphql?query={states(search:"A"){name}}'
```

