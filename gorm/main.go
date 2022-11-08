package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const dns = "host=localhost user=postgres password=abc123. dbname=learning port=5555 sslmode=disable TimeZone=Europe/Madrid"

func main() {
	configureDB(dns)

	getInstance().AutoMigrate(new(Food), new(Warehouse), new(User), new(Company), new(Library), new(Book))

	ctx, cl := context.WithTimeout(context.Background(), 10*time.Second)
	defer cl()

	DB := getInstanceWithCtx(ctx)

	var result []Library
	// DB.Model(&Library{}).Preload("Book").Find(&result) // same as the next
	DB.Preload("Book").Find(&result, "libraries.name = ?", "La nobel")

	buf, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf))
}

// _ = DB.Joins("Company").Find(&result)
// DB.Joins("Company", DB.Where(&Company{Name: "viewnext"})).Find(&result) // this doesn't work
// DB.Preload("Company").Joins("inner join companies on users.company_id = companies.id and companies.name like ?", "viewnext").Find(&result) // this does work
