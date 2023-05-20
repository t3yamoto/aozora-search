package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var dsn string
	flag.StringVar(&dsn, "d", "database.sqlite", "database")
	flag.Usage = func() {
		usage := "usage\n"
		fmt.Print(usage)
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch flag.Arg(0) {
	case "authors":
		err = showAuthors(db)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func showAuthors(db *sql.DB) error {
	rows, err := db.Query(`SELECT author_id, author FROM authors`)
	if err != nil {
		return err
	}
	for rows.Next() {
		var authorId string
		var author string
		err = rows.Scan(&authorId, &author)
		if err != nil {
			return err
		}
		fmt.Println(authorId, author)
	}
	return nil
}
