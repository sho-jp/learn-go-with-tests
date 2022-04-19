package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

const (
	dbType     = "mysql"
	dbSettings = "docker:docker@(127.0.0.1:3306)/db"
)

func getQuery(query string) int {
	db, err := sql.Open(dbType, dbSettings)

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	row := db.QueryRow(query)
	var score int

	row.Scan(&score)

	return score
}

func postQuery(query string) {
	db, err := sql.Open(dbType, dbSettings)

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	_, err = db.Exec(query)

	if err != nil {
		log.Println(err)
	}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	query := fmt.Sprintf("update players set score = score + 1 where name = %q", name)
	postQuery(query)
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	query := fmt.Sprintf("select score from players where name = %q", name)
	return getQuery(query)
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
