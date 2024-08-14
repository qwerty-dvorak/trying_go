package controller

import (
	"database/sql"
	"github.com/qwerty-dvorak/trying_go/schema"
)


func GetUsers(db *sql.DB) ([]schema.User, error) {
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