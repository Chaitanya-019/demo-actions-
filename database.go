package database

import (
	"database/sql"
	"fmt"
)

type Company struct {
	ID      string
	Name    string
	Creator string
}

// CompanyByID queries for the album with the specified ID.
func CompanyByID(id string, db *sql.DB) (Company, error) {
	// A Company to hold data from the returned row.
	var com Company
	row := db.QueryRow("SELECT * FROM company WHERE id = ?", id)
	if err := row.Scan(&com.ID, &com.Name, &com.Creator); err != nil {
		if err == sql.ErrNoRows {
			return com, fmt.Errorf("CompanyByID %d: no such Company", id)
		}
		return com, fmt.Errorf("CompanyByID %d: %v", id, err)
	}
	return com, nil
}
