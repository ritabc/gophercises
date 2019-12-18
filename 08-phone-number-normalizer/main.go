package main

import (
	"database/sql"
	"fmt"
	normalize "gophercises/08-phone-number-normalizer/numbers"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "rita"
	dbname = "phone_number_normalizer"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	deleteStmt := `DELETE FROM numbers;`
	res, err := db.Exec(deleteStmt)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted %d rows\n", count)

	numbersToInsert := []string{"1234567890", "123 456 7891", "(123) 456 7892", "(123) 456-7893", "123-456-7894", "123-456-7890", "1234567892", "(123)456-7892"}
	for _, raw := range numbersToInsert {
		sqlStatement := `INSERT INTO numbers (number) VALUES ($1) RETURNING id;`
		id := 0
		err = db.QueryRow(sqlStatement, raw).Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Inserted row with ID: %d\n", id)
	}

	selectStmt := `SELECT * FROM numbers`
	var id int
	var original string
	rows, err := db.Query(selectStmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &original)
		if err != nil {
			panic(err)
		}
		updateStmt := `UPDATE numbers SET number = $2 WHERE id = $1 RETURNING number;`
		var updated string
		err = db.QueryRow(updateStmt, id, normalize.Normalize(original)).Scan(&updated)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Updated number from %s to %s\n", original, updated)
	}

}
