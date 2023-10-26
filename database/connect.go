package database

import (
	"context"
	"log"
	"os"

	"github.com/mustafaakilll/ent_todo/ent"
)

// Connect function for connecting to database.
func Connect() *ent.Client {

	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
