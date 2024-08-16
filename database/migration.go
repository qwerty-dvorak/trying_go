package database

import (
    "database/sql"
    "log"
)

func RunMigration(db *sql.DB) {
    // Check if the users table exists
    var tableName string
    err := db.QueryRow(`
        SELECT table_name 
        FROM information_schema.tables 
        WHERE table_schema = 'public' AND table_name = 'users'
    `).Scan(&tableName)

    if err != nil && err != sql.ErrNoRows {
        log.Fatalf("Error checking for users table: %v", err)
    }

    // If the table does not exist, create it
    if tableName == "" {
        _, err := db.Exec(`
            CREATE TABLE users (
                id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                email TEXT UNIQUE NOT NULL,
                password TEXT NOT NULL
            )
        `)
        if err != nil {
            log.Fatalf("Error creating users table: %v", err)
        } else {
            log.Println("Users table created successfully.")
        }
    } else {
        log.Println("Users table already exists.")
    }
	print("Migration ran successfully")
}