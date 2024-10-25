package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/zipkero/sample-ent-go/ent"
	"github.com/zipkero/sample-ent-go/package/util"
	"log"
)

func Example_create() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer util.Close(client)

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	task1, err := client.Todo.Create().Save(ctx)
	if err != nil {
		log.Fatalf("failed creating todo: %v", err)
	}
	fmt.Println(task1)
	// Output:
	// Todo(id=1)
}
