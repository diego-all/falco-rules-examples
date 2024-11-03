package main

// Based on Mario Carrion Post:
// https://mariocarrion.com/2021/10/22/golang-software-architecture-security-databases-sql-injection-permissions.html

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib" // To initialize `pgx` DB driver
)

// Passing in an `id` like:
//
// "33ca99ce-f1d1-46c2-b26e-a0ff45011b18' OR ''='"
//
// The value "33ca..." is irrelevant because the `OR` is used instead.

// go run main.go -id="33ca99ce-f1d1-46c2-b26e-a0ff45011b18' OR ''='"

func main() {
	var id string

	flag.StringVar(&id, "id", "", "User ID to delete")

	flag.Parse()

	if id == "" {
		fmt.Println("id is blank, exiting")
		return
	}

	conn, err := newConn()
	if err != nil {
		log.Fatalln("Couldn't connect to DB", err)
	}

	//-

	query := fmt.Sprintf("DELETE FROM users WHERE id = '%s'", id)

	if _, err := conn.Exec(context.Background(), query); err != nil {
		log.Fatalln("Couldn't delete", err)
	}
}

func newConn() (*pgx.Conn, error) {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("user", "password"),
		Host:   "localhost:54329",
		Path:   "dbname",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	conn, err := pgx.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return conn, nil
}
