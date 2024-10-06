package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" 
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		log.Fatal(err)
	}
	
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to SQLite database!")
	return db
}

func CreateTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "code" TEXT,
        "original_link" TEXT
    );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Table created successfully!")
}

func AddLink(db *sql.DB, link string) error {
	var code string

	for {
		code = generateRandomCode(5)

		var exists bool
		err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM users WHERE code = ?)`, code).Scan(&exists)
		if err != nil {
			return err
		}
		if !exists {
			break 
		}
	}

	if _, err := db.Exec(`INSERT INTO users (code, original_link) VALUES (?, ?)`, code, link); err != nil {
		return err
	}

	log.Println("New link added to the database successfully!")
	return nil
}

func GetLinkFromDB(db *sql.DB, code string) (string, error) {
	querySQL := `SELECT original_link FROM users WHERE code = ?`

	var link string
	err := db.QueryRow(querySQL, code).Scan(&link)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	log.Println("Link retrieved from the database successfully!")
	return link, nil
}

func GetCodeFromDB(db *sql.DB, original_link string) (string, error) {
	querySQL := `SELECT code FROM users WHERE original_link = ?`

	var link string
	err := db.QueryRow(querySQL, original_link).Scan(&link)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	log.Println("Link retrieved from the database successfully!")
	return link, nil
}
