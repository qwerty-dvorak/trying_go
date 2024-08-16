package controller

import (
	"database/sql"
	"github.com/qwerty-dvorak/trying_go/schema"
)


func ReadUsers(db *sql.DB) ([]schema.User, error) {
	query:=`SELECT * FROM users`
	rows, err:=db.Query(query)
	if err != nil {
		print(err)
	}
	defer rows.Close()
	var users []schema.User
	for rows.Next() {
		var user schema.User
		err:=rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			print(err)
		}
		users=append(users, user)
	}
	return users, nil
}

func ReadUser(db *sql.DB, email string) (schema.User, error) {
	var user schema.User
	query:=`SELECT * FROM users WHERE email=$1`
	err:=db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
        return user, err
    }
    return user, nil
}

func CreateUser(db *sql.DB, user schema.CreateUser) (schema.User, error) {
	var useradd schema.User
	existingUser, err:=ReadUser(db, user.Email)
	if err==nil {
        return existingUser, err
    } else if err != nil {
        return useradd, err
    }
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
    err = db.QueryRow(query, user.Email, user.Password).Scan(&useradd.ID)
    if err != nil {
        return useradd, err
    }
	useradd.Email=user.Email
	useradd.Password=user.Password
	return useradd, nil
}