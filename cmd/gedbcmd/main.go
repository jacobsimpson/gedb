package main

import (
	"database/sql"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()

	filename := "./TestOpenDatabase.gedb"
	db, err := sql.Open("gedb", filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open the database %q: %v\n", filename, err)
		os.Exit(1)
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to execute query: %v\n")
		os.Exit(1)
	}
	for rows.Next() {
		fmt.Printf("%v\n", rows.Scan())
	}
}
