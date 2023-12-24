package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	connStr := "postgresql://postgres:csd1993@localhost/gs23_paymentdb?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Connected...")

	insertStmt := `insert into users (username, name, phone) values ($1, $2, $3)`
	result, err := db.Exec(insertStmt, "deniz", "Deniz Karan", "+905325158012")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	count, _ := result.RowsAffected()
	fmt.Printf("%d rows affected!...\n", count)
	err = db.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
