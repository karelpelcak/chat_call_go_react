package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
)

func Migration() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
    id SERIAL PRIMARY KEY,
    filename TEXT UNIQUE NOT NULL,
    applied_at TIMESTAMP DEFAULT now()
)`)

	if err != nil {
		log.Fatal(err)
	}

	files, err := filepath.Glob("migrations/*.sql")
	if err != nil {
		log.Fatal(err)
	}

	sort.Strings(files)
	for _, file := range files {
		filename := filepath.Base(file)

		var exists bool
		err := DB.QueryRow("SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE filename = $1)", filename).Scan(&exists)
		if err != nil {
			log.Fatal("ahoj", err)
		}

		if exists {
			fmt.Println("Skipping:", filename)
			continue
		}

		sqlBytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		queries := string(sqlBytes)

		_, err = DB.Exec(queries)
		if err != nil {
			log.Fatalf("Error running migration %s: %v", filename, err)
		}

		_, err = DB.Exec("INSERT INTO schema_migrations (filename) VALUES ($1)", filename)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Applied:", filename)
	}
}
