package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hasanozgan/frodao/drivers/postgres"
	"github.com/hasanozgan/frodao/examples/dal"
	"github.com/hasanozgan/frodao/nullable"
)

func main() {
	ctx := context.Background()
	// DSN Format: postgresql://user:pass@localhost:5432/db?sslmode=disable
	DSN := os.Getenv("DSN")
	if err := postgres.Connect(DSN); err != nil {
		log.Fatalf("DB Connection failed %s", DSN)
	}
	defer postgres.Close()

	userDAO := dal.NewUserDAO()
	record, err := userDAO.Create(ctx, &dal.UserTable{
		Username: "netologist",
		Password: "s3Cr3TP@55w0Rd",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created: %v\n", record)

	record, err = userDAO.FindByID(ctx, record.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected: %v\n", record)

	record.Address = nullable.New("47 Old Brompton Road, London, SW7 3JP")
	err = userDAO.Update(ctx, record)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated: %v\n", record)

	record, err = userDAO.FindByUsername(ctx, record.Username)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected by custom query: %v\n", record)

	err = userDAO.Delete(ctx, record.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Record Deleted: %v\n", record)

	record, err = userDAO.FindByID(ctx, record.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Find Record Deleted: %v\n", record)
}
