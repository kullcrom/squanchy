package db

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kullcrom/squanchy/types"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

//Connect will initialize a connection to the database.
func Connect() (db *sql.DB) {
	dbPassword, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		panic("ERROR: DB_PASSWORD not found")
	}
	db, err := sql.Open("mysql", "root:"+dbPassword+"@/tit?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//CreateUser creates a new User in the users table.
func CreateUser(email string, password string) error {
	db := Connect()
	defer db.Close()

	checkUser, err := GetUserByEmail(email)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	} else if checkUser.ID != 0 {
		return errors.New("ERROR: User already exists")
	}

	insert, err := db.Prepare("INSERT INTO users (email, password, first_name, last_name, created_on, last_login) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	result, err := insert.Exec(email, hashedPassword, "", "", time.Now(), time.Now())
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	} else if rows <= 0 {
		return errors.New("ERROR: No rows were affected")
	}

	log.Printf("Creating user for %v... %v row(s) affected.", email, rows)
	return nil
}

//DeleteUserByEmail deletes the specified User from the users table.
func DeleteUserByEmail(email string) error {
	db := Connect()
	defer db.Close()

	userToDelete, err := GetUserByEmail(email)
	if err != nil {
		return err
	}

	delete, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		return err
	}

	result, err := delete.Exec(userToDelete.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	} else if rows <= 0 {
		return errors.New("ERROR: No rows were affected")
	}

	log.Printf("Deleting user %v... %v row(s) affected.", email, rows)
	return nil
}

//GetUserByID queries the database for the specific user based on the user's ID.
func GetUserByID(id int) (types.User, error) {
	db := Connect()
	defer db.Close()

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

	return user, nil
}

//GetUserByEmail queries the database for the specific user based on the user's email.
func GetUserByEmail(email string) (types.User, error) {
	db := Connect()
	defer db.Close()

	result := db.QueryRow("SELECT id, email, first_name, last_name FROM users WHERE email=?", email)

	user := types.User{}
	var id int
	var queryEmail, firstName, lastName string
	err := result.Scan(&id, &queryEmail, &firstName, &lastName)
	if err != nil {
		return user, err
	}

	user.ID = id
	user.Email = queryEmail
	user.FirstName = firstName
	user.LastName = lastName

	return user, nil
}

//GetUserPasswordByEmail ...
func GetUserPasswordByEmail(email string) (string, error) {
	db := Connect()
	defer db.Close()

	result := db.QueryRow("SELECT (password) FROM users WHERE email=?", email)

	var password string
	err := result.Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}

//GetAllUsers ...
func GetAllUsers() ([]types.User, error) {
	db := Connect()
	defer db.Close()

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

	return users, nil
}

//UpdateUser ...
func UpdateUser(user *types.User) error {
	db := Connect()
	defer db.Close()

	checkUser, err := GetUserByEmail(user.Email)
	if err != nil {
		return err
	}

	update, err := db.Prepare("UPDATE users SET email = ?, first_name = ?, last_name = ? WHERE id = ?")
	if err != nil {
		return err
	}

	result, err := update.Exec(user.Email, user.FirstName, user.LastName, checkUser.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	} else if rows <= 0 {
		return errors.New("ERROR: No rows were affected")
	}

	return nil
}
