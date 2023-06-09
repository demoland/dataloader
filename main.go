package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	Name           string
	UserID         string
	Address        string
	Phone          string
	UserAgent      string
	Company        string
	Email          string
	Team           string
	Location       string
	CreditCard     string
	SocialSecurity string
}

func main() {
	// Define flag arguments
	filePtr := flag.String("file", "", "Path to the users JSON file (required)")
	dbNamePtr := flag.String("database", "", "Database name (required)")
	hostPtr := flag.String("hostname", "", "Database host (required)")
	portPtr := flag.String("port", "", "Database port (required)")
	passwordPtr := flag.String("password", "", "Database password (required)")

	// Parse flag arguments
	flag.Parse()

	// Check if required flags are provided
	if *filePtr == "" || *dbNamePtr == "" || *hostPtr == "" || *portPtr == "" || *passwordPtr == "" {
		flag.PrintDefaults()
		log.Fatal("All flags are required.")
	}

	// Read users JSON file
	data, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to PostgreSQL
	connStr := fmt.Sprintf("user=postgres password=%s dbname=%s host=%s port=%s sslmode=disable", *passwordPtr, *dbNamePtr, *hostPtr, *portPtr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the users table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		userid UUID,
		address TEXT,
		phone TEXT,
		useragent TEXT,
		company TEXT,
		email TEXT,
		team TEXT,
		location TEXT,
		creditcard TEXT,
		socialsecurity TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Insert users into the database
	for _, user := range users {
		_, err := db.Exec(`INSERT INTO users (
			name,  userid, address, phone, useragent, company, email, team, location, creditcard, socialsecurity
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
			user.Name, user.UserID, user.Address, user.Phone, user.UserAgent, user.Company, user.Email, user.Team, user.Location, user.CreditCard, user.SocialSecurity)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Data loaded successfully.")
}
