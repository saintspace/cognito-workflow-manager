package main

import (
	"database/sql"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	lambda.Start(handlePostConfirmation)
}

func handlePostConfirmation(event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	fmt.Println("Event:", event)
	userEmail := event.Request.UserAttributes["email"]
	fmt.Println("User email:", userEmail)
	// Initialize the user in your datastore
	// err = initializeUserInDatastore(userEmail)
	// if err != nil {
	// 	return "", err
	// }

	return event, nil
}

func initializeUserInDatastore(userEmail string) error {
	db, err := connectToPlanetScaleDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Generate a UUID
	userID := uuid.New().String()

	// Replace 'users' and column names with your actual table and column names
	insertQuery := "INSERT INTO users (id, email) VALUES (?, ?)"
	_, err = db.Exec(insertQuery, userID, userEmail)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}

	return nil
}

func connectToPlanetScaleDB() (*sql.DB, error) {
	// Replace these with your actual PlanetScale connection details
	user := "your-username"
	password := "your-password"
	host := "your-host"
	port := "your-port"
	database := "your-database"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
