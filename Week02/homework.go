package homework

import (
	"database/sql"
	"errors"
)

# DAO layer
func GetItems() ([]interface{}, error){
	rows, err := db.QueryContext(ctx, "SELECT items FROM table")
	if err != nil {
		return nil, fmt.Errorf("querying from database: %w", err)
	}

	// ignore success case
}

# Logic layer
func() Do() error {
	items, err := GetItems()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// discard the error
			return nil
		} else {
			return err
		}
	}

	// ... ignore success case
}