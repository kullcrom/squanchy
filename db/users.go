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
		panic("DB_PASSWORD not found")
	}
	db, err := sql.Open("mysql", "root:"+dbPassword+"@/squanchy")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//CreateUser creates a new User in the users table.
func CreateUser(user types.User) (types.UserResponse, error) {
	db := Connect()

	newUser := *types.NewUserResponse(user.Email, user.FirstName, user.LastName)

	checkUser, err := GetUserByEmail(user.Email)
	if err != nil {
		return newUser, err
	}
	if checkUser.ID != 0 {
		existingUser := *types.NewUserResponse(checkUser.Email, checkUser.FirstName, checkUser.LastName)
		return existingUser, errors.New("User already exists")
	}

	insert, err := db.Prepare("INSERT INTO users (email, password, first_name, last_name) VALUES (?, ?, ?, ?)")
	if err != nil {
		return newUser, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return newUser, err
	}

	result, err := insert.Exec(user.Email, hashedPassword, user.FirstName, user.LastName)
	if err != nil {
		return newUser, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return newUser, err
	}
	log.Printf("Creating user for %v %v... %v row(s) affected.", user.FirstName, user.LastName, rows)

	defer db.Close()
	if rows <= 0 {
		return newUser, errors.New("An error may have occurred when inserting the user as no rows were affected")
	}
	return newUser, nil
}

//DeleteUser deletes the specified User from the users table.
func DeleteUser(user types.User) (types.UserResponse, error) {
	db := Connect()

	userToDelete, err := GetUserByEmail(user.Email)
	if err != nil {
		return types.UserResponse{
			Email:     userToDelete.Email,
			FirstName: userToDelete.FirstName,
			LastName:  userToDelete.LastName,
		}, err
	}

	if userToDelete.ID == 0 {
		return types.UserResponse{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}, errors.New("User does not exist and therefore cannot be deleted")
	}

	deletedUser := *types.NewUserResponse(userToDelete.Email, userToDelete.FirstName, userToDelete.LastName)

	delete, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		return deletedUser, err
	}

	result, err := delete.Exec(user.ID)
	if err != nil {
		return deletedUser, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return deletedUser, err
	}
	log.Printf("Deleting user %v %v... %v row(s) affected.", user.FirstName, user.LastName, rows)

	defer db.Close()
	if rows <= 0 {
		return deletedUser, errors.New("An error may have occurred when deleting the user, as no rows were affected")
	}
	return deletedUser, nil
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
func GetUserByEmail(email string) (types.User, error) {
	db := Connect()

	user := types.User{}
	user.Email = email

	result, err := db.Query("SELECT * FROM users WHERE email=?", email)
	if err != nil {
		return user, err
	}

	for result.Next() {
		var id int
		var email, firstName, lastName string
		err := result.Scan(&id, &email, &firstName, &lastName)
		if err != nil {
			return user, err
		}

		user.ID = id
		user.Email = email
		user.FirstName = firstName
		user.LastName = lastName
	}

	defer db.Close()
	return user, nil
}
