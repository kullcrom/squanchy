package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kullcrom/squanchy/types"
	"log"
	"os"
)

//Connect will initialize a connection to the database.
func Connect() (db *sql.DB) {
	dbPassword, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		panic("DB_PASSWORD not found")
	}
	db, err := sql.Open("mysql", "root:"+dbPassword+"@/squanchy")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//CreateUser creates a new User in the users table.
func CreateUser(user types.User) bool {
	db := Connect()

	checkUser := GetUserByEmail(user.Email)
	if checkUser.ID != 0 {
		return false
	}

	insert, err := db.Prepare("INSERT INTO users (email, first_name, last_name) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	result, err := insert.Exec(user.Email, user.FirstName, user.LastName)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Creating user for %v %v... %v row(s) affected.", user.FirstName, user.LastName, rows)

	defer db.Close()
	if rows <= 0 {
		return false
	}
	return true
}

//DeleteUser deletes the specified User from the users table.
func DeleteUser(user types.User) bool {
	db := Connect()

	delete, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	result, err := delete.Exec(user.ID)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleting user %v %v... %v row(s) affected.", user.FirstName, user.LastName, rows)

	defer db.Close()
	if rows <= 0 {
		return false
	}
	return true
}

//GetUserByID queries the database for the specific user based on the user's ID.
func GetUserByID(id int) types.User {
	db := Connect()

	result, err := db.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}

	user := types.User{}
	for result.Next() {
		var id int
		var email, firstName, lastName string
		err := result.Scan(&id, &email, &firstName, &lastName)
		if err != nil {
			log.Fatal(err)
		}

		user.ID = id
		user.Email = email
		user.FirstName = firstName
		user.LastName = lastName
	}

	log.Printf("Found user %v %v", user.FirstName, user.LastName)
	defer db.Close()
	return user
}

//GetUserByEmail queries the database for the specific user based on the user's email.
func GetUserByEmail(email string) types.User {
	db := Connect()

	result, err := db.Query("SELECT * FROM users WHERE email=?", email)
	if err != nil {
		log.Fatal(err)
	}

	user := types.User{}
	for result.Next() {
		var id int
		var email, firstName, lastName string
		err := result.Scan(&id, &email, &firstName, &lastName)
		if err != nil {
			log.Fatal(err)
		}

		user.ID = id
		user.Email = email
		user.FirstName = firstName
		user.LastName = lastName
	}

	defer db.Close()
	return user
}
