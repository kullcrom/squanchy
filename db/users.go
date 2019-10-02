package db

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kullcrom/squanchy/types"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

//Connect will initialize a connection to the database.
func Connect() (db *sql.DB) {
	dbPassword, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		panic("ERROR: DB_PASSWORD not found")
	}
	db, err := sql.Open("mysql", "root:"+dbPassword+"@/squanchy")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//CreateUser creates a new User in the users table.
func CreateUser(user types.User) (types.User, error) {
	db := Connect()

	checkUser, err := GetUserByEmail(user.Email)
	if err != nil {
		log.Println(err)
	}
	if checkUser.ID != 0 {
		return checkUser, errors.New("ERROR: User already exists")
	}

	insert, err := db.Prepare("INSERT INTO users (email, password, first_name, last_name) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	result, err := insert.Exec(user.Email, hashedPassword, user.FirstName, user.LastName)
	if err != nil {
		log.Println(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	log.Printf("Creating user for %v %v... %v row(s) affected.", user.FirstName, user.LastName, rows)

	newUser, err := GetUserByEmail(user.Email)
	if err != nil {
		log.Println(err)
	}

	defer db.Close()
	if rows <= 0 {
		return newUser, errors.New("ERROR: No rows were affected")
	}
	return newUser, nil
}

//DeleteUserByEmail deletes the specified User from the users table.
func DeleteUserByEmail(user types.User) (types.User, error) {
	db := Connect()

	userToDelete, err := GetUserByEmail(user.Email)
	if err != nil {
		return user, err
	}

	if userToDelete.ID == 0 {
		return user, errors.New("ERROR: User not found")
	}

	delete, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		return user, err
	}

	result, err := delete.Exec(user.ID)
	if err != nil {
		return user, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return user, err
	}
	log.Printf("Deleting user %v %v... %v row(s) affected.", user.FirstName, user.LastName, rows)

	defer db.Close()
	if rows <= 0 {
		return user, errors.New("ERROR: No rows were affected")
	}
	return user, nil
}

//GetUserByID queries the database for the specific user based on the user's ID.
func GetUserByID(id int) (types.User, error) {
	db := Connect()

	user := types.User{}
	result, err := db.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return user, err
	}

	for result.Next() {
		var id int
		var email, password, firstName, lastName string
		err := result.Scan(&id, &email, &password, &firstName, &lastName)
		if err != nil {
			return user, err
		}

		user.ID = id
		user.Email = email
		user.Password = password
		user.FirstName = firstName
		user.LastName = lastName
	}

	log.Printf("Found user %v %v", user.FirstName, user.LastName)
	defer db.Close()
	return user, nil
}

//GetUserByEmail queries the database for the specific user based on the user's email.
func GetUserByEmail(email string) (types.User, error) {
	db := Connect()

	result := db.QueryRow("SELECT * FROM users WHERE email=?", email)

	user := types.User{}
	var id int
	var queryEmail, password, firstName, lastName string
	err := result.Scan(&id, &queryEmail, &password, &firstName, &lastName)
	if err != nil {
		return user, err
	}

	user.ID = id
	user.Email = queryEmail
	user.Password = password
	user.FirstName = firstName
	user.LastName = lastName

	defer db.Close()
	return user, nil
}

//GetAllUsers ...
func GetAllUsers() ([]types.User, error) {
	db := Connect()

	var users []types.User

	result, err := db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}

	user := types.User{}
	for result.Next() {
		var id int
		var email, password, firstName, lastName string
		err := result.Scan(&id, &email, &password, &firstName, &lastName)
		if err != nil {
			return users, err
		}

		user.ID = id
		user.Email = email
		user.Password = password
		user.FirstName = firstName
		user.LastName = lastName
		users = append(users, user)
	}

	defer db.Close()
	return users, nil
}
