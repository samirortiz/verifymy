package users

import (
	"database/sql"
)

var db *sql.DB

// SET DATABASE CONNECTION
func SetDB(database *sql.DB) {
	db = database
}

// RETURN ALL USERS
func allUsers() ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.Id, &usr.Name, &usr.Age, &usr.Email, &usr.Password, &usr.Address); err != nil {
			return nil, err
		}
		users = append(users, usr)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// RETURN SINGLE USER BY ID
func userByID(id string) (User, error) {
	var user User
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	if err := row.Scan(&user.Id, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}
	return user, nil
}

// CREATE USER
func createUser(user User) (User, error) {
	stmt, _ := db.Prepare("INSERT INTO users(name, age, address, email, password) VALUES (?, ?, ?, ?, ?)")
	res, err := stmt.Exec(user.Name, user.Age, user.Address, user.Email, user.Password)
	if err != nil {
		return user, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return user, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", lastId)
	if err := row.Scan(&user.Id, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}
	return user, nil
}

// UPDATE USER BY ID
func updateUser(userId string, user User) (User, error) {
	stmt, _ := db.Prepare("UPDATE users set name = ?, age = ?, address = ?, email = ?, password = ? WHERE id = ?")
	res, err := stmt.Exec(user.Name, user.Age, user.Address, user.Email, user.Password, userId)
	if err != nil {
		return user, err
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		return user, err
	}
	if rowAffected != 1 {
		return user, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", userId)
	if err := row.Scan(&user.Id, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}
	return user, nil
}

// DELETE USER BY ID
func deleteUser(id int) (int, error) {
	res, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	rowAffected, err := res.RowsAffected()
	if rowAffected == 0 {
		return 0, err
	}
	return id, nil
}
