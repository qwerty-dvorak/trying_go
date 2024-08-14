package main

import(
	"github.com/qwerty-dvorak/trying_go/database"
)

func main() {
	db, err := database.NewSession()
	if err != nil {
		print(err)
	}
	print(db,"\n")
    println("Hello, Go!")
}