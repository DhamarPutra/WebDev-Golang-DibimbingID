package models

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Portofolio struct {
	ID          int
	Title       string
	Description string
	Image       string
}

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	// Buat tabel jika belum ada
	createTable := `CREATE TABLE IF NOT EXISTS portofolios (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		image TEXT
	);`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllPortofolios() ([]Portofolio, error) {
	rows, err := DB.Query("SELECT id, title, description, image FROM portofolios")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var portofolios []Portofolio
	for rows.Next() {
		var p Portofolio
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Image)
		if err != nil {
			return nil, err
		}
		portofolios = append(portofolios, p)
	}
	return portofolios, nil
}

func CreatePortofolio(title, description, image string) error {
	_, err := DB.Exec("INSERT INTO portofolios (title, description, image) VALUES (?, ?, ?)", title, description, image)
	return err
}

func DeletePortofolio(id int) error {
	_, err := DB.Exec("DELETE FROM portofolios WHERE id=?", id)
	return err
}
