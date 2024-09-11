package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// load env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

    // Open the CSV file
    file, err := os.Open("Placement_Data_Full_Class.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Read the CSV file
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // Establish a connection to the PostgreSQL database
	user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_DATABASE")
    host := os.Getenv("POSTGRES_HOST")
    port := os.Getenv("POSTGRES_PORT")
    sslmode := "disable"

    connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s", user, dbname, password, host, port, sslmode)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Prepare an SQL statement for inserting data
    stmt, err := db.Prepare(`
        INSERT INTO placement_data (sl_no, gender, ssc_p, ssc_b, hsc_p, hsc_b, hsc_s, degree_p, degree_t, workex, etest_p, specialisation, mba_p, status, salary)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
    `)
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    // Iterate over the rows of the CSV file and execute the prepared statement for each row
    for _, record := range records[1:] { // Skip header row
        _, err := stmt.Exec(record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8], record[9], record[10], record[11], record[12], record[13], record[14])
        if err != nil {
            log.Fatal(err)
        }
    }
}