package config

import "database/sql"

func InitDB(migrations []string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "gophisher.db")
	if err != nil {
		return nil, err
	}

	for _, migration := range migrations {
		_, err = db.Exec(migration)
		if err != nil {
			return nil, err
		}

	}

	return db, nil
}
