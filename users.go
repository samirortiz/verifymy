package main

import (
	"database/sql"
	"net/url"
)

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

func createUser(formData url.Values) (User, error) {
	var user User
	stmt, _ := db.Prepare("INSERT INTO users(name, age, address, email, password) VALUES (?, ?, ?, ?, ?)")
	res, err := stmt.Exec(formData.Get("name"), formData.Get("age"), formData.Get("address"), formData.Get("email"), formData.Get("password"))
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

func updateUser(userId string, formData url.Values) (User, error) {
	var user User
	stmt, _ := db.Prepare("UPDATE users set name = ?, age = ?, address = ?, email = ?, password = ? WHERE id = ?")
	res, err := stmt.Exec(formData.Get("name"), formData.Get("age"), formData.Get("address"), formData.Get("email"), formData.Get("password"), userId)
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
