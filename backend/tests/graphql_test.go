package tests

import (
	"context"
	"select-location/internal/config"
	"select-location/internal/db"
	"select-location/internal/models"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestStateSearch(t *testing.T) {
	client, err := db.ConnectDB(config.MongoURI)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	tests := []struct {
		name     string
		search   string
		expected int
	}{
		{
			name:     "Search A returns multiple states",
			search:   "A",
			expected: 5, // Updated: Alabama, Alaska, American Samoa, Arizona, Arkansas
		},
		{
			name:     "Search New returns multiple states",
			search:   "New",
			expected: 4, // New Hampshire, New Jersey, New Mexico, New York
		},
		{
			name:     "Search Z returns no states",
			search:   "Z",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := bson.D{}

			if tt.search != "" {
				filter = bson.D{{
					"name",
					bson.D{{
						"$regex",
						"^" + tt.search,
					}, {
						"$options",
						"i",
					}},
				}}
			}

			var states []models.State
			cursor, err := client.Database(config.DBName).Collection("states").Find(context.TODO(), filter)
			if err != nil {
				t.Fatal(err)
			}
			defer cursor.Close(context.TODO())

			if err = cursor.All(context.TODO(), &states); err != nil {
				t.Fatal(err)
			}

			if len(states) != tt.expected {
				t.Errorf("expected %d states, got %d", tt.expected, len(states))
			}
		})
	}
}

// More tests can be added, like TestCaseSensitivity, TestEmptySearch and TestDatabaseConnection if time permits.
