package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/jacobsimpson/gedb"
	flag "github.com/spf13/pflag"
)

const VERSION = "0.0.1"

func execName() string {
	return filepath.Base(os.Args[0])
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] FILENAME [SQL]\n", execName())
		fmt.Fprintf(os.Stderr, `
Command line utility for interacting with GEDB databases.

FILENAME is the name of the GEDB database to query. A new database is created
if the file does not already exist.

`)
		fmt.Fprintln(os.Stderr, flag.CommandLine.FlagUsagesWrapped(80))
	}
}

func main() {
	var version bool
	var verbose int

	flag.CountVarP(&verbose, "verbose", "v",
		"increase output for debugging purposes")
	flag.BoolVar(&version, "version", version,
		"output version information and exit")
	flag.Parse()

	if version {
		fmt.Printf("%s %s\n", execName(), VERSION)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "A database file to use must be specified.")
		os.Exit(1)
	}
	filename := args[0]
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "A SQL query to execute must be specified.")
		os.Exit(1)
	}
	query := args[1]

	db, err := sql.Open("gedb", filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open the database %q: %v\n", filename, err)
		os.Exit(1)
	}

	rows, err := db.Query(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to execute query: %v\n")
		os.Exit(1)
	}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to list the columns for the result set: %v\n", err)
		os.Exit(1)
	}
	for _, c := range columns {
		fmt.Printf("%s\n", c)
	}
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to get data: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Row = %d, %q\n", id, name)
	}
}
