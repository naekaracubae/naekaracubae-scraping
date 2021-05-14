package etc

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "gobbyisntfree.ccttcm80dlu1.ap-northeast-2.rds.amazonaws.com"
	database = "GobbyIsntFree"
	user     = ""
	password = ""
)

func MakeConnection() {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	CheckErr(err)
	defer db.Close()

	err = db.Ping()
	CheckErr(err)
	fmt.Println("Successfully created connection to database.")

	var (
		name  string
		email string
	)

	rows, err := db.Query("SELECT name, email from subscribers;")
	CheckErr(err)
	defer rows.Close()
	fmt.Println("Reading data:")
	for rows.Next() {
		err := rows.Scan(&name, &email)
		CheckErr(err)
		fmt.Printf("Data row = (%s, %s)\n", name, email)
	}

	err = rows.Err()
	CheckErr(err)
	fmt.Println("Done.")

}

func GetSubscribers() {

}

func GetSenders() {

}
