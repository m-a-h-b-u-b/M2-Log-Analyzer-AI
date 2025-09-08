package storage

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(path string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", path+"?_busy_timeout=5000&_journal_mode=WAL")
	if err != nil {
		return nil, err
	}
	store := &SQLiteStore{db: db}
	if err := store.migrate(); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *SQLiteStore) migrate() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			received_at DATETIME,
			raw TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS anomalies (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			log_id INTEGER,
			detected_at DATETIME,
			detector TEXT,
			score REAL,
			details TEXT,
			FOREIGN KEY(log_id) REFERENCES logs(id)
		);`,
	}
	for _, q := range queries {
		if _, err := s.db.Exec(q); err != nil {
			return err
		}
	}
	return nil
}

func (s *SQLiteStore) InsertLog(raw string) (int64, error) {
	res, err := s.db.Exec(`INSERT INTO logs(received_at, raw) VALUES (?, ?)`, time.Now().UTC(), raw)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (s *SQLiteStore) InsertAnomaly(logID int64, detector string, score float64, details string) (int64, error) {
	res, err := s.db.Exec(`INSERT INTO anomalies(log_id, detected_at, detector, score, details) VALUES (?, ?, ?, ?, ?)`,
		logID, time.Now().UTC(), detector, score, details)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// Example: fetch recent logs for training
func (s *SQLiteStore) FetchLogs(limit int) ([]string, error) {
	rows, err := s.db.Query(`SELECT raw FROM logs ORDER BY received_at DESC LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []string{}
	for rows.Next() {
		var raw string
		if err := rows.Scan(&raw); err != nil {
			return nil, err
		}
		out = append(out, raw)
	}
	return out, nil
}

func (s *SQLiteStore) Close() error {
	log.Println("Closing sqlite db")
	return s.db.Close()
}
