package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	const (
		DB_USER     = "psimoes"
		DB_PASSWORD = "password"
		DB_NAME     = "gotest"
	)
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM PRODUCT")
	if err != nil {
		fmt.Println(err)
	}

	arr := []string{}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("asf")
		}
		arr = append(arr, name)
	}
	file, err := os.OpenFile("/home/psimoes/Desktop/teste.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("asf")

	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(arr)
	if err != nil {
		fmt.Errorf("LULU")
	}
	fmt.Println("vim-go")
}
