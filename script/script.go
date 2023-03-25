package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/kenta-ja8/clean-arch/pkg/config"
	"github.com/kenta-ja8/clean-arch/pkg/util"
	_ "github.com/mattn/go-sqlite3"
)

func setup(db *sql.DB) {

	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS "user" (
      "id" CHAR(36) PRIMARY KEY,
      "name" VARCHAR(255),
      "birthday" DATETIME
    )`,
	)
	if err != nil {
		log.Panicln(err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "book" (
      "id" CHAR(36) PRIMARY KEY,
      "title" VARCHAR(255)
    )`,
	)
	if err != nil {
		log.Panicln(err)
	}
}
func clean(db *sql.DB) {
	_, err := db.Exec(
		`DROP TABLE IF EXISTS "user"`,
	)
	if err != nil {
		log.Panicln(err)
	}
	_, err = db.Exec(
		`DROP TABLE IF EXISTS "book"`,
	)
	if err != nil {
		log.Panicln(err)
	}
}

func test(db *sql.DB) {
	log.Println(util.NewUUIDv4())
	_, err := db.Exec(
		`INSERT INTO user(id, name, birthday) VALUES (?, ?, ?)`,
		util.NewUUIDv4(),
		"name",
		"2023-01-01T00:00:00Z",
	)
	if err != nil {
		log.Panicln("failed to insert", err)
	}

	rows, err := db.Query("SELECT id, name, birthday FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	type User struct {
		Id       string
		Name     string
		Birthday string
	}

	var users []*User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Birthday); err != nil {
			log.Fatal(err)
		}
		users = append(users, &user)
	}
	for _, user := range users {
		log.Println(user)
	}

}

func main() {
	flag.Parse()
	flags := flag.Args()
	if len(flags) == 0 {
		log.Panicln("must has arg")
	}
	config := config.NewConfig()

	db, err := sql.Open(config.DBDriver, config.DBName)
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	switch flags[0] {
	case "setup":
		setup(db)
		log.Println("finish setup")
	case "test":
		test(db)
		log.Println("finish test")
	case "clean":
		clean(db)
		log.Println("finish clean")
	}
}
