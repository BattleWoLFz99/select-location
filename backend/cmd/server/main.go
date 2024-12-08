package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"select-location/internal/config"
	"select-location/internal/db"
	gql "select-location/internal/graphql"

	"github.com/graphql-go/graphql"
	"github.com/rs/cors"
)

func main() {
	client, err := db.ConnectDB(config.MongoURI)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	schema, err := gql.CreateSchema(client)
	if err != nil {
		panic(err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	http.Handle("/graphql", c.Handler(handler))

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
