package controller

import (
	"database/sql"
	"github.com/qwerty-dvorak/trying_go/schema"
)


func ReadUsers(db *sql.DB) (schema.JsonReturn) {
	query:=`SELECT * FROM users`
	rows, err:=db.Query(query)
	if err != nil {
		return schema.JsonReturn{Status: "error in query", Message: err.Error(), Data: nil}
	}
	defer rows.Close()
	var users []schema.User
	for rows.Next() {
		var user schema.User
		err:=rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return schema.JsonReturn{Status: "error in row", Message: err.Error(), Data: nil}
		}
		users=append(users, user)
	}
	return schema.JsonReturn{Status: "success", Message: "Users found", Data: users}
}

func ReadUser(db *sql.DB, email string) (schema.JsonReturn) {
	var user schema.User
	query:=`SELECT * FROM users WHERE email=$1`
	err:=db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return schema.JsonReturn{Status: "error", Message: "User not found", Data: nil}
	} else if err != nil {
        return schema.JsonReturn{Status: "error in query", Message: err.Error(), Data: nil}
    }
    return schema.JsonReturn{Status: "success", Message: "User found", Data: user}
}

func CreateUser(db *sql.DB, user schema.CreateUser) (schema.JsonReturn) {
	var useradd schema.User
	existingUser:=ReadUser(db, user.Email)
	if existingUser.Status=="success" {
		if existingUser.Data.(schema.User).Password==user.Password {
		return schema.JsonReturn{Status: "error", Message: "User already exists", Data: nil}
		}	else {
			return schema.JsonReturn{Status: "error", Message: "Wrong password", Data: nil}
		}
	}
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
    err := db.QueryRow(query, user.Email, user.Password).Scan(&useradd.ID)
    if err != nil {
        return schema.JsonReturn{Status: "error in inserting", Message: err.Error(), Data: nil}
    }
	useradd.Email=user.Email
	useradd.Password=user.Password
	return schema.JsonReturn{Status: "success", Message: "User added", Data: useradd}
}