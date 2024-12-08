# select-location

## Features
- Import a json file that contains locations
- Typeahead suggestion
- Interactive map display using Google Maps API
- GraphQL API for efficient data querying
- Responsive design for various screen sizes

This project utilized Go, vue.js, docker, and GraphQL

## Installation
Please make sure Go, Docker, vue3 is properly installed

### Backend
#### Install dependencies
```sh
go get github.com/graphql-go/graphql
go get go.mongodb.org/mongo-driver/mongo
go get github.com/rs/cors
```

#### Data seed
Redirect to ./backend
```sh
docker-compose up -d
go run seed.go
```
#### Verify data seed
```sh
docker exec -it YOUR_CONTAINER_NAME mongosh
use us_states db.states.find() 
```
#### Start GraphQL server
```sh
go run main.go
```
#### Test API
```sh
http://localhost:8080/graphql?query={states(search:"A"){name}}
```

### Frontend
Apply for a Google API Key and enable the "Maps JavaScript API" and "Geocoding API". Paste the Google API Key into the YOUR_API_KEY section of the file located at ./frontend/src/components/SelectState.vue.
```sh
npm install -g pnpm
pnpm install
pnpm dev
```
Then go to http://localhost:3000/

