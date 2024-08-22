package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migrateDB(config map[string]string) {
	dbUser := config["db_username"]
	dbPassword := config["db_password"]
	dbHost := config["db_host"]
	dbName := config["db_name"]
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

    db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	dir, _ := os.Getwd()
    m, err := migrate.NewWithDatabaseInstance(
        fmt.Sprintf("file://%s/db_migrations", dir),
        "mysql", 
        driver,
    )
	if err != nil {
		log.Fatal(err)
	}

    err = m.Up()
	if err != nil {
		log.Fatal(err)
	}
}