package v1

import (
	"database/sql"
	"fmt"
	"log"
)

type UserAccount struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func RegistrationData(user UserAccount) {
	connStr := "user=username dbname=mydb sslmode=disable password=mypassword"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if it does not exist
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username TEXT NOT NULL,
            email TEXT NOT NULL,
            password TEXT NOT NULL
        );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Insert user data into the database
	sqlStatement := `
        INSERT INTO users (username, email, password)
        VALUES ($1, $2, $3)
        RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New record ID is: %d\n", id)
}
