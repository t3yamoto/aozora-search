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
	case "titles":
		if flag.NArg() != 2 {
			flag.Usage()
			os.Exit(2)
		}
		err = showTitles(db, flag.Arg(1))
	case "content":
		if flag.NArg() != 3 {
			flag.Usage()
			os.Exit(3)
		}
		err = showContent(db, flag.Arg(1), flag.Arg(2))
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
	defer rows.Close()

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

func showTitles(db *sql.DB, authorID string) error {
	rows, err := db.Query(`SELECT title_id, title FROM contents WHERE author_id = ?`, authorID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var titleId string
		var title string
		err = rows.Scan(&titleId, &title)
		if err != nil {
			return err
		}
		fmt.Printf("% 5s %s\n", titleId, title)
	}
	return nil
}

func showContent(db *sql.DB, authorID string, titleID string) error {
	var content string
	err := db.QueryRow(`SELECT content FROM contents WHERE author_id = ? AND title_id = ?`, authorID, titleID).Scan(&content)
	if err != nil {
		return err
	}
	fmt.Println(content)
	return nil
}
