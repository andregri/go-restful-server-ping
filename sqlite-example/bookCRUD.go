package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	// Create db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal("open error:", err)
	} else {
		log.Println(db)
	}

	// Create table
	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY,
			isbn INTEGER,
			author VARCHAR(64),
			name VARCHAR(64) NULL
		)
	`)
	if err != nil {
		log.Fatal("Error in creating table")
	} else {
		log.Println("Successfully created table books")
	}
	statement.Exec()

	// Create record
	createStatement, err := db.Prepare(`
		INSERT INTO books (name, author, isbn)
		VALUES (?, ?, ?)
	`)
	if err != nil {
		log.Println("insert error:", err)
	} else {
		createStatement.Exec("A tale of two cities", "Charles Dickens", 140430547)
		log.Println("Inserted the book into database!")
	}

	// Read records
	rows, err := db.Query(`
		SELECT id, name, author FROM  books
	`)
	if err != nil {
		log.Println("read error:", err)
	} else {
		var tempBook Book
		for rows.Next() {
			rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
			log.Printf("ID: %d, Book: %q, Author: %q",
				tempBook.id, tempBook.name, tempBook.author)
		}
	}

	// Update records
	updateStatement, err := db.Prepare(`
		UPDATE books SET name=? WHERE id=?
	`)
	if err != nil {
		log.Println("update error:", err)
	} else {
		updateStatement.Exec("The Tale of Two Cities", 1)
		log.Println("Successfully updated the book in database")
	}

	// Delete records
	deleteStatement, err := db.Prepare(`
		DELETE FROM books WHERE id=?
	`)
	if err != nil {
		log.Println("delete error:", err)
	} else {
		deleteStatement.Exec(1)
		log.Println("Successfully deleted the book in database")
	}
}
