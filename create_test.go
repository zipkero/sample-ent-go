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

	task1, err := client.Todo.Create().SetText("add graphql example").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating todo: %v", err)
	}
	fmt.Printf("%d: %q\n", task1.ID, task1.Text)
	task2, err := client.Todo.Create().SetText("add tracking example").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating todo: %v", err)
	}
	fmt.Printf("%d: %q\n", task2.ID, task2.Text)
	// Output:
	// 1: "add graphql example"
	// 2: "add tracking example"
}
