package db

import (
	"database/sql"
	"gnomiki/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	conn *sql.DB
}

func New(dbPath string) (*DB, error) {
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	
	db := &DB{conn: conn}
	if err := db.createTable(); err != nil {
		return nil, err
	}
	
	return db, nil
}

func (db *DB) createTable() error {
	query := `CREATE TABLE IF NOT EXISTS gnomiki (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	)`
	_, err := db.conn.Exec(query)
	return err
}

func (db *DB) GetAll() ([]models.Gnomik, error) {
	rows, err := db.conn.Query("SELECT id, name, age FROM gnomiki")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var gnomiki []models.Gnomik
	for rows.Next() {
		var g models.Gnomik
		if err := rows.Scan(&g.ID, &g.Name, &g.Age); err != nil {
			return nil, err
		}
		gnomiki = append(gnomiki, g)
	}
	return gnomiki, nil
}

func (db *DB) Create(g models.Gnomik) error {
	_, err := db.conn.Exec("INSERT INTO gnomiki (name, age) VALUES (?, ?)", g.Name, g.Age)
	return err
}