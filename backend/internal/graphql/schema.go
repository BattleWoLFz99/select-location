package graphql

import (
	"context"
	"select-location/internal/models"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateSchema(client *mongo.Client) (graphql.Schema, error) {
	fields := graphql.Fields{
		"states": &graphql.Field{
			Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
				Name: "State",
				Fields: graphql.Fields{
					"name": &graphql.Field{Type: graphql.String},
				},
			})),
			Args: graphql.FieldConfigArgument{
				"search": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				collection := client.Database("us_states").Collection("states")
				filter := bson.D{}

				if search, ok := p.Args["search"].(string); ok && search != "" {
					filter = bson.D{{
						"name",
						bson.D{{
							"$regex",
							"^" + search,
						}, {
							"$options",
							"i",
						}},
					}}
				}

				var states []models.State
				cursor, err := collection.Find(context.TODO(), filter)
				if err != nil {
					return nil, err
				}
				defer cursor.Close(context.TODO())
				if err = cursor.All(context.TODO(), &states); err != nil {
					return nil, err
				}
				return states, nil
			},
		},
	}

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Query",
			Fields: fields,
		}),
	})
}
