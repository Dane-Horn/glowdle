package wordle

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type db struct {
	*sql.DB
}

type word struct {
	word             string
	allowed_solution bool
}

func connect() *db {
	database, err := sql.Open("sqlite3", `/home/dane/glowdle/wordle.db?nolock=1`)
	if err != nil {
		log.Fatal(err)
	}
	database.SetMaxOpenConns(1)
	return &db{database}
}

func scanWord(row *sql.Row) (word, error) {
	var word word
	err := row.Scan(&word.word, &word.allowed_solution)
	if err != nil {
		return word, err
	}
	return word, nil
}

func (db *db) getRandomSolution() string {
	var word word
	row := db.QueryRow(`
    select * from allowed_solutions
    order by random()
    limit 1;
  `)

	word, err := scanWord(row)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ToUpper(word.word)
}

func (db *db) validGuess(guess string) bool {
	var count int
	row := db.QueryRow(`
    select count(*) from words
    where word = ?
  `, strings.ToLower(guess))
	row.Scan(&count)
	return count > 0
}
